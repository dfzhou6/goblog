package article

import (
	"github.com/dfzhou6/goblog/app/models"
	"github.com/dfzhou6/goblog/app/models/user"
	"github.com/dfzhou6/goblog/pkg/route"
	"github.com/dfzhou6/goblog/pkg/types"
)

type Article struct {
	models.BaseModel
	Title      string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body       string `gorm:"type:longtext;not null;" valid:"body"`
	UserID     uint64 `gorm:"not null;index"`
	User       user.User
	CategoryID uint64 `gorm:"not null;default:4;index"`
}

func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
