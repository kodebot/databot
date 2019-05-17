package controllers

import (
	"fmt"

	"github.com/revel/revel"
)

// Article controller
type Article struct {
	*revel.Controller
}

// List returns list of articles
func (c Article) List() revel.Result {
	queries := c.Request.GetQuery()

	fmt.Printf("%+v", queries)
	return c.RenderJSON(struct{}{})
}
