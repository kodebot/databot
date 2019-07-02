package controllers

import (
	"github.com/kodebot/newsfeed/conf"
	"github.com/revel/revel"
)

// Category controller
type Category struct {
	*revel.Controller
}

// List returns all available categories
func (c Category) List() revel.Result {
	return c.RenderJSON(conf.AppSettings.ArticleCategory)
}
