package services

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

func loadFeedConfig(t *testing.T, filePath string) models.FeedConfig {
	if filePath == "" {
		filePath = "../conf/feed_parsing_config.toml"
	}
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile(filePath, &feedConfig)
	if err != nil {
		t.Fatalf("error when loading feed config: %s\n", err.Error())
	}
	return feedConfig
}

func getFeedConfigItemByName(t *testing.T, feedConfigName string, filePath string) *models.FeedConfigItem {
	feedConfigs := loadFeedConfig(t, filePath)

	for _, feedConfigItem := range feedConfigs.Feed {
		if feedConfigItem.Name == feedConfigName {
			return &feedConfigItem
		}
	}

	t.Fatalf("Feed config item with name %s not found", feedConfigName)
	return nil
}

type SingleCreateArticleTest struct {
	feedConfigName  string
	testFeedFile    string
	expectedArticle models.Article
}

var createArticleTestCases = []SingleCreateArticleTest{
	{
		"dinamalar_politics",
		"./test_data/feeds/dinamalar_politics_20190603.xml",
		models.Article{
			Title:         "கட்சி தாவுகிறாரா திவ்யா ஸ்பந்தனா",
			ShortContent:  "",
			PublishedDate: time.Date(2019, 6, 2, 16, 3, 0, 0, time.UTC), // Sun, 02 Jun 2019 21:33:00 +0530
			SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2289413",
			Categories:    []string{"politics"},
			ThumbImageURL: "http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2289413_150_100.jpg",
			Source:        "dinamalar"}},
	{
		"dinamalar_general_pot1",
		"./test_data/feeds/dinamalar_general_pot1_20190605.xml",
		models.Article{
			Title:         "பறவைகளுக்கு விலாசம் சொன்னது யார்?",
			ShortContent:  "",
			PublishedDate: time.Date(2019, 6, 4, 18, 4, 0, 0, time.UTC), // Tue, 04 Jun 2019 23:34:00 +0530
			SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2290872",
			Categories:    []string{"general"},
			ThumbImageURL: "http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2290872_150_100.jpg",
			Source:        "dinamalar"}},
	{
		"dinamalar_incidents_sam1",
		"./test_data/feeds/dinamalar_incidents_sam1_20190605.xml",
		models.Article{
			Title:         "கேரள மாணவருக்கு, 'நிபா'தனி வார்டில் தீவிர சிகிச்சை",
			ShortContent:  "",
			PublishedDate: time.Date(2019, 6, 4, 18, 27, 0, 0, time.UTC), // Tue, 04 Jun 2019 23:34:00 +0530
			SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2290888",
			Categories:    []string{"incidents"},
			ThumbImageURL: "http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2290888_150_100.jpg",
			Source:        "dinamalar"}},
	{
		"dinamalar_justice_kut1",
		"./test_data/feeds/dinamalar_justice_kut1_20190605_1.xml",
		models.Article{
			Title:         "புதுச்சேரி முதல்வருக்கு, 'நோட்டீஸ்'",
			ShortContent:  "",
			PublishedDate: time.Date(2019, 6, 4, 18, 44, 0, 0, time.UTC), // Wed, 05 Jun 2019 00:14:00 +0530
			SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2290904",
			Categories:    []string{"justice"},
			ThumbImageURL: "",
			Source:        "dinamalar"}},
	{
		"dinamalar_justice_kut1",
		"./test_data/feeds/dinamalar_justice_kut1_20190605_2.xml",
		models.Article{
			Title:         "ஒட்டுமொத்த கல்வி முறையையும் சீரமையுங்கள்: உச்ச நீதிமன்றம் உத்தரவு",
			ShortContent:  "",
			PublishedDate: time.Date(2019, 6, 4, 20, 13, 0, 0, time.UTC), // Wed, 05 Jun 2019 01:43:00 +0530
			SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2291013",
			Categories:    []string{"justice"},
			ThumbImageURL: "http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2291013_150_100.jpg",
			Source:        "dinamalar"}}}

func TestCreateArticle(t *testing.T) {
	for _, testCase := range createArticleTestCases {
		t.Run(fmt.Sprintf("CreateArticle tests for %s, %s", testCase.feedConfigName, testCase.testFeedFile), func(t *testing.T) {
			feedConfigName := testCase.feedConfigName
			testFeedFile := testCase.testFeedFile
			feedConfig := getFeedConfigItemByName(t, feedConfigName, "")

			// ** test setup start **
			if feedConfig == nil {
				t.Fatalf("Unable to load feed config with name %s", feedConfigName)
			}

			testFeedXMLBytes, err := ioutil.ReadFile(testFeedFile)

			if err != nil {
				t.Fatalf("Unable to load test data from %s, %s", testFeedFile, err.Error())
			}

			// todo: test fixing illegal chars

			parsedFeeds := ParseFeed(*feedConfig, string(testFeedXMLBytes))

			if parsedFeeds == nil {
				t.Fatalf("Unable to parse feed data from file %s with feed config name %s", testFeedFile, feedConfigName)
			}

			// ** test setup end **

			preConditionsPass := runCreateArticlePreConditionTests(t, parsedFeeds)
			if preConditionsPass {
				runCreateArticleTests(t, testCase, parsedFeeds, feedConfig)
			}
		})
	}
}

func runCreateArticlePreConditionTests(outerT *testing.T, parsedFeeds []*gofeed.Item) bool {
	return outerT.Run("<<PRE-CONDITIONS>>", func(t *testing.T) {
		// feed item count
		expected := 1
		actual := len(parsedFeeds)
		if actual != expected {
			// make outer testing to fail
			t.FailNow()
			outerT.Fatalf("expected **%d** parsed feeds but found **%d**", expected, actual)
		}
	})
}

func runCreateArticleTests(
	t *testing.T,
	testCase SingleCreateArticleTest,
	parsedFeeds []*gofeed.Item,
	feedConfig *models.FeedConfigItem) {
	testArticles := CreateArticles(parsedFeeds, *feedConfig)
	actualArticle := testArticles[0]
	expectedArticle := testCase.expectedArticle

	t.Run("<<BUSINESS TESTS>>", func(t *testing.T) {
		if actualArticle.Title != expectedArticle.Title {
			t.Errorf("Title >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.Title, actualArticle.Title)
		}

		if actualArticle.ShortContent != expectedArticle.ShortContent {
			t.Errorf("ShortContent >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.ShortContent, actualArticle.ShortContent)
		}

		if !actualArticle.PublishedDate.Equal(expectedArticle.PublishedDate) {
			t.Errorf("PublishedDate >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.PublishedDate, actualArticle.PublishedDate)
		}

		if len(actualArticle.Categories) != len(expectedArticle.Categories) {
			t.Errorf("Categories count >> EXPECTED: **%d** ACTUAL: **%d**", len(expectedArticle.Categories), len(actualArticle.Categories))
		}

		for _, expectedCategory := range expectedArticle.Categories {
			found := false
			for _, actualCategory := range actualArticle.Categories {
				if expectedCategory == actualCategory {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Category missing >> EXPECTED: **%s** but not found", expectedCategory)
			}
		}

		if actualArticle.ThumbImageURL != expectedArticle.ThumbImageURL {
			t.Errorf("ThumbImageURL >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.ThumbImageURL, actualArticle.ThumbImageURL)
		}

		if actualArticle.SourceURL != expectedArticle.SourceURL {
			t.Errorf("SourceURL >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.SourceURL, actualArticle.SourceURL)
		}

		if actualArticle.Source != expectedArticle.Source {
			t.Errorf("Source >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.Source, actualArticle.Source)
		}

		now := time.Now()
		twoSecBefore := now.Add(time.Second * -2)
		if !((actualArticle.CreatedAt.Equal(now) || actualArticle.CreatedAt.Before(now)) &&
			actualArticle.CreatedAt.After(twoSecBefore)) {
			t.Errorf("CreatedAt EXPECTED: between **%s** and **%s** ACTUAL: **%s**", twoSecBefore, now, actualArticle.CreatedAt)
		}
	})
}
