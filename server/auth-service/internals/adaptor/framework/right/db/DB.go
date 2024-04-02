package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/models"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"
)

type Adapter struct {
	config ports.ConfigPort
	DB     *pgx.Conn
}

func Initialize(cfg ports.ConfigPort) *Adapter {
	return &Adapter{
		config: cfg,
		DB:     nil,
	}
}

func (dbAd *Adapter) DBInit() error {
	cfg := dbAd.config.GetConfig()
	env := cfg.Env
	host := env["db_host"]
	port := env["db_port"]
	user := env["db_user"]
	password := env["db_password"]
	dbname := env["db_database"]
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	var maxRetries int
	var err error
	var db *pgx.Conn
	var retryInterval int = 2
	for {
		maxRetries++
		db, err = pgx.Connect(context.Background(), dbURL)
		if err == nil || maxRetries > 10 {
			break
		}
		cfg.Logger.Printf("Failed to connect to PostgreSQL: %v", err)
		cfg.Logger.Printf("Retrying attempt %d in %vsec(s)...", maxRetries, retryInterval)
		time.Sleep(time.Duration(retryInterval) * time.Second)
	}
	if err != nil {
		return fmt.Errorf("exceeded maximum connection retries. unable to connect to postgres")
	}
	_, err = db.Exec(context.Background(), CreateUserTableQuery())
	if err != nil {
		return fmt.Errorf("error creating the table: %v", err.Error())
	}
	dbAd.DB = db
	return nil
}

func (dbAd *Adapter) GetDB() *pgx.Conn {
	return dbAd.DB
}

func (dbAd *Adapter) InsertUser(user models.User) (models.User, error) {
	query := InsertUserQuery()
	row := dbAd.DB.QueryRow(context.Background(),
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
	)
	var newUser models.User
	err := row.Scan(
		&newUser.ID,
		&newUser.FirstName,
		&newUser.LastName,
		&newUser.Email,
		&newUser.Active,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (dbAd *Adapter) InsertUsers(users []models.User) ([]models.User, error) {
	tx, err := dbAd.DB.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())
	query := InsertUserQuery()
	var insertedUsers []models.User
	for _, user := range users {
		var newUser models.User
		err = tx.QueryRow(context.Background(),
			query,
			user.FirstName,
			user.LastName,
			user.Email,
			user.Password,
			time.Now(),
			time.Now(),
		).Scan(
			&newUser.ID,
			&newUser.FirstName,
			&newUser.LastName,
			&newUser.Email,
			&newUser.Active,
			&newUser.CreatedAt,
			&newUser.UpdatedAt,
		)
		if err != nil {
			tx.Rollback(context.Background())
			return nil, err
		}
		insertedUsers = append(insertedUsers, newUser)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return insertedUsers, nil
}

func (dbAd *Adapter) GetUserByEmail(email string) (models.User, error) {
	query := GetUserByEmailQuery()
	var user models.User
	row := dbAd.DB.QueryRow(context.Background(),
		query,
		email,
	)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (dbAd *Adapter) GetUserByID(userID uint) (models.User, error) {
	query := GetUserByIdQuery()
	var user models.User
	row := dbAd.DB.QueryRow(context.Background(),
		query,
		userID,
	)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (dbAd *Adapter) DeleteUserByID(userID uint) error {
	query := DeleteUserByIdQuery()
	commandTag, err := dbAd.DB.Exec(context.Background(), query, userID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("user with ID %d not found or not deleted", userID)
	}
	return nil
}
