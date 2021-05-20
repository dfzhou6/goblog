package config

import "github.com/dfzhou6/goblog/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{
		"perpage":   10,
		"url_query": "p",
	})
}
