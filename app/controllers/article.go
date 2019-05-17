package controllers

import (
	"fmt"
	"strconv"

	"github.com/golang/glog"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kodebot/newsfeed/models"

	"github.com/kodebot/newsfeed/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/revel/revel"
)

// Article controller
type Article struct {
	*revel.Controller
}

// List returns list of articles
func (c Article) List() revel.Result {
	page := c.Params.Query.Get("page")
	if page == "" {
		page = "1"
	}

	fmt.Printf("%s\n", page)

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		glog.Warningf("parsing page number failed %s. setting to 1\n", page)
		pageInt = 1
	}

	articleCollection, err := data.GetCollection("articles")

	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		c.Response.Status = 500
		return c.RenderText("Internal error")
	}

	findOptions := options.Find().SetSkip((pageInt - 1) * 20).SetLimit(20).SetProjection(bson.M{"_id": 1, "title": 1, "publisheddate": 1, "categories": 1, "thumbimageurl": 1, "sourceurl": 1, "source": 1})

	result, _ := data.Find(articleCollection, bson.M{}, func(cursor *mongo.Cursor) interface{} {
		var article models.ArticleMinimal
		err := cursor.Decode(&article)
		if err != nil {
			glog.Warningf("error while decoding %+v", cursor.Current)
		}
		return article

	}, findOptions)

	return c.RenderJSON(result)
}
