package ports

import (
	"github.com/jackc/pgx/v5"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"
)

type DBPort interface {
	DBInit() error
	GetDB() *pgx.Conn
	InsertUser(models.User) (models.User, error)
	InsertUsers([]models.User) ([]models.User, error)
	GetUserByEmail(string) (models.User, error)
	GetUserByID(uint) (models.User, error)
	DeleteUserByID(uint) error
}
