package db

func CreateUserTableQuery() string {
	return `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		active BOOLEAN,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`
}

func AlterTableIDSequence() string {
	return `ALTER SEQUENCE users_id_seq restart 1000 OWNED BY users.id`
}

func InsertUserQuery() string {
	return `INSERT INTO users (first_name, last_name, email, password, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, true, $5, $6)
		RETURNING id, first_name, last_name, email, active, created_at, updated_at`
}

func InsertInitUserQuery() string {
	return `INSERT INTO users (first_name, last_name, email, password, active, user_type, branch_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, true, $5, $6, $7, $8)
		RETURNING id, first_name, last_name, email, active, user_type, branch_id, created_at, updated_at`
}

func GetUserByEmailQuery() string {
	return `SELECT id, first_name, last_name, email, password, active, user_type, branch_id, created_at, updated_at
		FROM USERS
		WHERE email=$1`
}

func GetUserByIdQuery() string {
	return `SELECT id, first_name, last_name, email, password, active, user_type, branch_id, created_at, updated_at
		FROM users
		WHERE id=$1`
}

func DeleteUserByIdQuery() string {
	return `DELETE FROM users WHERE id=$1`
}
