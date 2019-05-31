package controllers

import (
	"github.com/kodebot/newsfeed/conf"
	"github.com/kodebot/newsfeed/models"
	"github.com/revel/revel"
)

type Category struct {
	*revel.Controller
}

func (c Category) List() revel.Result {
	var articleCategories []models.ArticleCategory

	for _, category := range conf.AppSettings.ArticleCategory {
		if category.IsPublic == true {
			articleCategories = append(articleCategories, category)
		}
	}
	return c.RenderJSON(articleCategories)
}
