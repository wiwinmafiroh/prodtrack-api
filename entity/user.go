package entity

import (
	"os"
	"prodtrack-api/pkg/errs"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AccessRole string

const (
	AdminRole AccessRole = "admin"
	UserRole  AccessRole = "user"
)

var secretKey = os.Getenv("SECRET_KEY")

var invalidTokenErr = errs.NewUnauthenticatedError("Invalid token")

type User struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      AccessRole `json:"role"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
}

func (u *User) HashPassword() errs.ErrorResponse {
	salt := 8
	password := []byte(u.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) GenerateToken() string {
	claims := u.claimsToken()

	return u.signToken(claims)
}

func (u *User) claimsToken() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"role":  u.Role,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, _ := token.SignedString([]byte(secretKey))

	return stringToken
}

func (u *User) ValidateToken(bearerToken string) errs.ErrorResponse {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")
	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, "Bearer ")
	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	stringToken := splitToken[1]

	token, err := u.parseToken(stringToken)
	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) parseToken(stringToken string) (*jwt.Token, errs.ErrorResponse) {
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claims jwt.MapClaims) errs.ErrorResponse {
	if id, ok := claims["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = uint(id)
	}

	if email, ok := claims["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = email
	}

	if role, ok := claims["role"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Role = AccessRole(role)
	}

	return nil
}
