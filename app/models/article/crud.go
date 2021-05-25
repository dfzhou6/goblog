package article

import (
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
