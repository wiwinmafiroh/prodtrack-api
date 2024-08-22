package user_postgres

import (
	"database/sql"
	"prodtrack-api/entity"
	"prodtrack-api/pkg/errs"
	"prodtrack-api/repository/user_repository"

	"github.com/lib/pq"
)

const (
	createUserQuery = `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4);
	`

	getUserByEmailQuery = `
		SELECT id, name, email, password, role FROM users
		WHERE email = $1;
	`
)

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) user_repository.UserRepository {
	return &userPostgres{
		db: db,
	}
}

func (u *userPostgres) CreateUser(userEntity entity.User) errs.ErrorResponse {
	_, err := u.db.Exec(createUserQuery, userEntity.Name, userEntity.Email, userEntity.Password, userEntity.Role)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok && pqError.Code == "23505" {
			return errs.NewBadRequestError("Email already exists")
		}

		return errs.NewInternalServerError("Failed to create new user")
	}

	return nil
}

func (u *userPostgres) GetUserByEmail(userEmail string) (*entity.User, errs.ErrorResponse) {
	row := u.db.QueryRow(getUserByEmailQuery, userEmail)

	var retrievedUser entity.User

	err := row.Scan(&retrievedUser.Id, &retrievedUser.Name, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		}

		return nil, errs.NewInternalServerError("Failed to retrieve user data")
	}

	return &retrievedUser, nil
}
