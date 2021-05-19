package view

import (
	"github.com/dfzhou6/goblog/pkg/auth"
	"github.com/dfzhou6/goblog/pkg/flash"
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/dfzhou6/goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	var err error
	data["isLogined"] = auth.Check()
	data["flash"] = flash.All()
	data["Categories"], err = nil, nil

	allFiles := getTemplateFiles(tplFiles...)
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(allFiles...)
	logger.LogError(err)

	tmpl.ExecuteTemplate(w, name, data)
}

func getTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	return append(layoutFiles, tplFiles...)
}
