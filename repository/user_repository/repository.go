package user_repository

import (
	"prodtrack-api/entity"
	"prodtrack-api/pkg/errs"
)

type UserRepository interface {
	CreateUser(userEntity entity.User) errs.ErrorResponse
	GetUserByEmail(userEmail string) (*entity.User, errs.ErrorResponse)
}
