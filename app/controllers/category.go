package controllers

import (
	"github.com/kodebot/newsfeed/conf"
	"github.com/revel/revel"
)

type Category struct {
	*revel.Controller
}

func (c Category) List() revel.Result {
	return c.RenderJSON(conf.AppSettings.ArticleCategory)
}
