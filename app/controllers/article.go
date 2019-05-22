package controllers

import (
	"fmt"
	"strconv"

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

// List returns list of articles
func (c Article) List() revel.Result {
	page := c.Params.Query.Get("page")
	category := c.Params.Query.Get("category")

	if page == "" {
		page = "1"
	}

	if category == "" {
		category = "0"
	}

	fmt.Printf("%s\n", page)

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		glog.Warningf("parsing page number failed %s. setting to 1\n", page)
		pageInt = 1
	}

	categoryInt, err := strconv.ParseInt(category, 10, 64)

	if err != nil {
		glog.Warningf("parsing category number failed %s. setting to 0\n", category)
		categoryInt = 0
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

	/*
		available categories
		general = 0
		politics = 1
		incident = 0
		tamilnadu = 0
		delhi = 0
		cinema = 2
		sports = 4
		world = 3
		business = 0

	*/

	var filter interface{}

	switch categoryInt {
	case 0:
		filter = bson.M{"categories": bson.M{"$nin": []string{"politics", "cinema", "world", "sports"}}}
	case 1:
		filter = bson.M{"categories": bson.M{"$in": []string{"politics"}}}
	case 2:
		filter = bson.M{"categories": bson.M{"$in": []string{"cinema"}}}
	case 3:
		filter = bson.M{"categories": bson.M{"$in": []string{"world"}}}
	case 4:
		filter = bson.M{"categories": bson.M{"$in": []string{"sports"}}}
	}

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
