package category

import (
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/dfzhou6/goblog/pkg/model"
	"github.com/dfzhou6/goblog/pkg/types"
)

func (c *Category) Create() (err error) {
	if err = model.DB.Create(c).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func All() (cs []Category, err error) {
	if err = model.DB.Find(&cs).Error; err != nil {
		return cs, err
	}
	return cs, nil
}

func Get(idStr string) (c Category, err error) {
	id := types.StringToInt(idStr)
	if err = model.DB.First(&c, id).Error; err != nil {
		return c, err
	}
	return c, nil
}
