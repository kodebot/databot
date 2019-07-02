package controllers

import (
	"github.com/kodebot/newsfeed/conf"
	"github.com/revel/revel"
)

// Source controller
type Source struct {
	*revel.Controller
}

// List returns all news sources available
func (c Source) List() revel.Result {
	return c.RenderJSON(conf.AppSettings.ArticleSource)
}
