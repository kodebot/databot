package scraper

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractGuidedConent(t *testing.T) {
	content := extractGuidedContent(
		"https://tamil.oneindia.com/news/delhi/abhinandhan-s-moushtache-was-a-talking-in-parliament-355077.html",
		"url",
		[]string{"#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt"},
		[]string{"#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > nav", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > h1", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > div.io-category", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > div.author-detail.clearfix", "#d_social_media", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > p.mat_promo", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > div.city-article-list-wrap", "#container > section > div > div.leftpanel > article > div > div.oiMiddleMain > div.oi-article-lt > div.io-author", "#notification-articleblock", ".deepLinkPromo"},
		"body > div.midbg.float-fix > div.left-box-container.float-fix > div > div.articlesub-cat-mid-rp > div > div:nth-child(4) > div.article-left-sub-cat1 > div > span > img")

	ioutil.WriteFile("./temp.html", []byte(content), os.ModePerm)
}
