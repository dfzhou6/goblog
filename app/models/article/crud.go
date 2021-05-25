package article

import (
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/dfzhou6/goblog/pkg/model"
	"github.com/dfzhou6/goblog/pkg/pagination"
	"github.com/dfzhou6/goblog/pkg/route"
	"github.com/dfzhou6/goblog/pkg/types"
	"net/http"
)

func Get(idStr string) (a Article, err error) {
	id := types.StringToInt(idStr)
	if err = model.DB.Preload("User").First(&a, id).Error; err != nil {
		return a, err
	}
	return a, nil
}

func GetAll(r *http.Request, perPage int) (cs []Article, data pagination.ViewData, err error) {
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)
	viewData := _pager.Paging()
	_pager.Results(&cs)
	return cs, viewData, nil
}

func (a *Article) Create() (err error) {
	if err = model.DB.Create(&a).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (a *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (a *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func GetByUserID(uid string) (as []Article, err error) {
	if err = model.DB.Where("user_id = ?", uid).Preload("User").Find(&as).Error; err != nil {
		return as, err
	}
	return as, nil
}

func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)
	viewDta := _pager.Paging()
	var articles []Article
	_pager.Results(&articles)
	return articles, viewDta, nil
}
