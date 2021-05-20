package user

import (
	"github.com/dfzhou6/goblog/pkg/password"
	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
}
