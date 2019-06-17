package rssatom

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/articles"
	datapkg "github.com/kodebot/newsfeed/data"
	field "github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collect returns collected fields from the given data using given field collector settings
func Collect(data string, fieldsInfo []field.Info) []map[string]interface{} {
	feeds := readFromXML(data)
	var records []map[string]interface{}
	for _, item := range feeds {
		//hack:

		itemLink := item.Link

		for _, fi := range fieldsInfo {
			if fi.Name == "SourceUrl" {
				if s := fi.CollectorInfo.Parameters["source"]; s != nil {
					sourceFieldName := s.(string)
					sourceField := reflect.Indirect(reflect.ValueOf(item)).FieldByName(sourceFieldName)
					if sourceField.IsValid() {
						itemLink = sourceField.String()
						break
					}
				}
			}
		}

		var existing articles.Article
		articlesCollection, err := datapkg.GetCollection("articles")
		err = datapkg.FindOne(articlesCollection, bson.M{"sourceurl": itemLink}, &existing)

		if err != nil && err != mongo.ErrNoDocuments {
			continue
		}

		if existing.SourceURL != "" {
			glog.Infoln("item already found, not adding this to the store...")
			continue
		}

		record := map[string]interface{}{}
		for _, fieldInfo := range fieldsInfo {
			record[fieldInfo.Name] = field.Create(item, fieldInfo)
		}
		records = append(records, record)
	}
	return records
}
