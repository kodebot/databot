package controllers

import (
	"github.com/kodebot/newsfeed/conf"
	"github.com/revel/revel"
)

type Source struct {
	*revel.Controller
}

func (c Source) List() revel.Result {
	return c.RenderJSON(conf.AppSettings.ArticleSource)
}
