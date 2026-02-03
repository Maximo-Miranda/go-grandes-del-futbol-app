package auth

import (
	"github.com/goravel/framework/contracts/http"

	"grandesdelfutbol/app/models"
)

type LoginResult struct {
	Success bool
	Token   string
	User    *models.User
	Error   string
}

type RegisterResult struct {
	Success bool
	User    *models.User
	Error   string
}

type AuthServiceInterface interface {
	Login(ctx http.Context, email, password, ip string) LoginResult
	Register(name, email, password string) RegisterResult
	Logout(ctx http.Context, userID uint) error
	GenerateToken(ctx http.Context, user *models.User) (string, error)
}
