package transformers

import "github.com/kodebot/newsfeed/datafeed/scraper"

func scrape(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	if html, found := params["scrappedHtml"]; found {
		return scraper.Scrape(html.(string), params)
	}

	return scraper.Scrape(val.(string), params)
}
