package route

import (
	"github.com/dfzhou6/goblog/pkg/config"
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return config.GetString("app.url") + url.String()
}

func GetRouteVariable(paramName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[paramName]
}
