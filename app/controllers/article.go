package controllers

import (
	"strconv"
	"strings"

	"github.com/kodebot/newsfeed/conf"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

// Get return article by id
func (c Article) Get(id string) revel.Result {
	articleCollection, err := data.GetCollection("articles")
	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		c.Response.Status = 500
		return c.RenderText("Internal error")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		glog.Warningf("error while parsing id %s error: %s", id, err.Error())
		c.Response.SetStatus(500)
		c.RenderText("Internal server error")
	}

	filter := bson.M{"_id": objectID}
	var result models.ArticleContent

	err = data.FindOne(articleCollection, filter, &result)
	if err != nil {
		glog.Warningf("error while getting article by id %s error: %s", id, err.Error())
	}
	return c.RenderJSON(result)
}

// RedirectToArticle redirects user to the original aritcle
func (c Article) RedirectToArticle(id string) revel.Result {
	articleCollection, err := data.GetCollection("articles")
	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		c.Response.Status = 500
		return c.RenderText("Internal error")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		glog.Warningf("error while parsing id %s error: %s", id, err.Error())
		c.Response.SetStatus(500)
		c.RenderText("Internal server error")
	}

	filter := bson.M{"_id": objectID}
	var result models.ArticleContent

	err = data.FindOne(articleCollection, filter, &result)
	if err != nil {
		glog.Warningf("error while getting article by id %s error: %s", id, err.Error())
	}
	return c.Redirect(result.SourceURL)
}

// List returns list of articles
func (c Article) List() revel.Result {
	page := c.Params.Query.Get("page")
	category := c.Params.Query.Get("category")
	sources := c.Params.Query.Get("sources")

	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		glog.Warningf("parsing page number failed %s. setting to 1\n", page)
		pageInt = 1
	}

	if category == "" {
		category = "0"
	}

	categoryInt, err := strconv.ParseInt(category, 10, 64)

	if err != nil {
		glog.Warningf("parsing category number failed %s. setting to 0\n", category)
		categoryInt = 0
	}

	var sourcesInt []int64
	for _, val := range strings.Split(strings.TrimSpace(sources), ",") {
		sourceInt, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			glog.Warningf("parsing source number failed %s. skipping...\n", val)
			continue
		} else {
			sourcesInt = append(sourcesInt, sourceInt)
		}
	}

	articleCollection, err := data.GetCollection("articles")

	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		c.Response.Status = 500
		return c.RenderText("Internal error")
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"publisheddate": -1})
	findOptions.SetSkip((pageInt - 1) * 20)
	findOptions.SetLimit(20)
	findOptions.SetProjection(bson.M{"_id": 1, "title": 1, "publisheddate": 1, "thumbimageurl": 1, "source": 1})

	var categoriesToFilter []string
	for _, category := range conf.AppSettings.ArticleCategory {
		if category.ID == int(categoryInt) {
			categoriesToFilter = append(categoriesToFilter, category.Category)
		}

		if category.ID == 0 { // general category - add all non public ones
			for _, cat := range conf.AppSettings.ArticleCategory {
				if cat.IsPublic != true {
					categoriesToFilter = append(categoriesToFilter, cat.Category)
				}
			}

		}
	}

	var sourcesToFilter []string
	for _, source := range conf.AppSettings.ArticleSource {
		for _, requestedSourceID := range sourcesInt {
			if source.ID == int(requestedSourceID) {
				sourcesToFilter = append(sourcesToFilter, source.Source)
			}
		}
	}

	if len(sourcesToFilter) == 0 {
		glog.Warningf("no sources specified, using all feeds")
		for _, source := range conf.AppSettings.ArticleSource {
			sourcesToFilter = append(sourcesToFilter, source.Source)
		}
	}

	filter := bson.M{"categories": bson.M{"$in": categoriesToFilter}, "source": bson.M{"$in": sourcesToFilter}}

	result, _ := data.Find(articleCollection, filter, func(cursor *mongo.Cursor) interface{} {
		var article models.ArticleMinimal
		err := cursor.Decode(&article)
		if err != nil {
			glog.Warningf("error while decoding %+v", cursor.Current)
		}
		return article

	}, findOptions)

	return c.RenderJSON(result)
}
