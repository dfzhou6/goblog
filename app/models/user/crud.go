package user

import (
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/dfzhou6/goblog/pkg/model"
	"github.com/dfzhou6/goblog/pkg/types"
)

func (u *User) Create() (err error) {
	if err = model.DB.Create(u).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func GetByEmail(email string) (u User, err error) {
	if err = model.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func Get(idStr string) (u User, err error) {
	id := types.StringToInt(idStr)
	if err = model.DB.First(&u, id).Error; err != nil {
		return u, err
	}
	return u, nil
}
