package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type User struct {
	orm.Model
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password string `gorm:"column:password;not null" json:"-"`
	Role     string `gorm:"column:role;default:player;not null" json:"role"`

	SignInCount     int        `gorm:"column:sign_in_count;default:0;not null" json:"sign_in_count"`
	CurrentSignInAt *time.Time `gorm:"column:current_sign_in_at" json:"-"`
	LastSignInAt    *time.Time `gorm:"column:last_sign_in_at" json:"-"`
	CurrentSignInIP *string    `gorm:"column:current_sign_in_ip" json:"-"`
	LastSignInIP    *string    `gorm:"column:last_sign_in_ip" json:"-"`
	FailedAttempts  int        `gorm:"column:failed_attempts;default:0;not null" json:"-"`
	LockedAt        *time.Time `gorm:"column:locked_at" json:"-"`
}

func (*User) TableName() string {
	return "users"
}

const lockoutDuration = 30 * time.Minute

func (u *User) IsLocked() bool {
	if u.LockedAt == nil {
		return false
	}
	if time.Since(*u.LockedAt) > lockoutDuration {
		return false
	}
	return true
}

func (u *User) Unlock() {
	u.LockedAt = nil
	u.FailedAttempts = 0
}

func (u *User) IncrementFailedAttempts() {
	u.FailedAttempts++
}

func (u *User) ResetFailedAttempts() {
	u.FailedAttempts = 0
}

func (u *User) ShouldLock() bool {
	return u.FailedAttempts >= 5
}

func (u *User) Lock() {
	now := time.Now()
	u.LockedAt = &now
}

func (u *User) TrackSignIn(ip string) {
	now := time.Now()
	u.LastSignInAt = u.CurrentSignInAt
	u.LastSignInIP = u.CurrentSignInIP
	u.CurrentSignInAt = &now
	u.CurrentSignInIP = &ip
	u.SignInCount++
	u.ResetFailedAttempts()
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) IsOrganizer() bool {
	return u.Role == "organizer"
}

func (u *User) HasRole(roles ...string) bool {
	for _, r := range roles {
		if u.Role == r {
			return true
		}
	}
	return false
}

// Ensure User implements Authenticatable for facades.Auth
func (u *User) GetKey() string {
	return ""
}

// GetPlayer returns the player profile associated with this user, if any
func (u *User) GetPlayer() (*Player, error) {
	var player Player
	if err := facades.Orm().Query().Where("user_id = ?", u.ID).First(&player); err != nil {
		return nil, err
	}
	if player.ID == 0 {
		return nil, nil
	}
	return &player, nil
}
