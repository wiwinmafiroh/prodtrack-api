package handler

import (
	"prodtrack-api/dto"
	"prodtrack-api/pkg/errs"
	"prodtrack-api/pkg/helpers"
	"prodtrack-api/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// UserRegister godoc
// @ID user-register
// @Summary Register a new user
// @Description Register a new user by providing the necessary information.
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.UserRegisterRequest true "User registration data. The role can be either 'admin' or 'user'."
// @Success 201 {object} dto.UserRegisterResponse
// @Router /users/register [post]
func (u *userHandler) UserRegister(ctx *gin.Context) {
	err := helpers.CheckContentType(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	var userRequest dto.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		errBindJSON := errs.NewUnprocessableEntityError("Invalid request body")

		ctx.AbortWithStatusJSON(errBindJSON.StatusCode(), errBindJSON)
		return
	}

	result, err := u.userService.UserRegister(userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// UserLogin godoc
// @ID user-login
// @Summary Log in a user
// @Description Log in a user by providing the necessary credentials.
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.UserLoginRequest true "User login data including email and password."
// @Success 200 {object} dto.UserLoginResponse
// @Router /users/login [post]
func (u *userHandler) UserLogin(ctx *gin.Context) {
	err := helpers.CheckContentType(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	var userRequest dto.UserLoginRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		errBindJSON := errs.NewUnprocessableEntityError("Invalid request body")

		ctx.AbortWithStatusJSON(errBindJSON.StatusCode(), errBindJSON)
		return
	}

	result, err := u.userService.UserLogin(userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
