package auth

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/models"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx http.Context, email, password, ip string) LoginResult {
	var user models.User
	if err := facades.Orm().Query().Where("email = ?", strings.ToLower(email)).First(&user); err != nil {
		return LoginResult{Success: false, Error: "Credenciales inválidas"}
	}

	// Auto-unlock if lockout period has expired
	if user.LockedAt != nil && !user.IsLocked() {
		user.Unlock()
		facades.Orm().Query().Save(&user)
	}

	if user.IsLocked() {
		return LoginResult{Success: false, Error: "Tu cuenta está bloqueada temporalmente. Intenta de nuevo en 30 minutos."}
	}

	if !facades.Hash().Check(password, user.Password) {
		user.IncrementFailedAttempts()
		if user.ShouldLock() {
			user.Lock()
		}
		facades.Orm().Query().Save(&user)
		return LoginResult{Success: false, Error: "Credenciales inválidas"}
	}

	user.TrackSignIn(ip)
	facades.Orm().Query().Save(&user)

	token, err := s.GenerateToken(ctx, &user)
	if err != nil {
		facades.Log().Error("Failed to generate JWT token: " + err.Error())
		return LoginResult{Success: false, Error: "Error al iniciar sesión"}
	}

	return LoginResult{Success: true, Token: token, User: &user}
}

func (s *AuthService) Register(name, email, password string) RegisterResult {
	count, _ := facades.Orm().Query().Model(&models.User{}).Where("email = ?", strings.ToLower(email)).Count()
	if count > 0 {
		// Return success to prevent user enumeration — same message regardless
		return RegisterResult{Success: true, User: nil}
	}

	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		return RegisterResult{Success: false, Error: "Error al crear la cuenta"}
	}

	user := models.User{
		Name:     name,
		Email:    strings.ToLower(email),
		Password: hashedPassword,
		Role:     "player",
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return RegisterResult{Success: false, Error: "Error al crear la cuenta"}
	}

	return RegisterResult{Success: true, User: &user}
}

func (s *AuthService) Logout(ctx http.Context, userID uint) error {
	return facades.Auth(ctx).Logout()
}

func (s *AuthService) GenerateToken(ctx http.Context, user *models.User) (string, error) {
	return facades.Auth(ctx).Login(user)
}
