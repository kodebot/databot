package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kodebot/newsfeed/logger"
)

// extractGuidedContent returns relevant content.
func extractGuidedContent(
	source string,
	sourceType string,
	focusSelectors []string,
	blacklistedSelectors []string,
	imgFallbackSelector string) string {

	// todo: update to detect the source type automatically
	var document *goquery.Document
	var err error

	if sourceType == "url" {
		if strings.HasPrefix(source, "//") {
			source = "http:" + source
		}
		document, err = goquery.NewDocument(source)
	}

	if sourceType == "html" {
		document, err = goquery.NewDocumentFromReader(strings.NewReader(source))
	}

	if err != nil {
		logger.Errorf("unable to create html document from %s ** error: %s ** sourceType: %s", source, err.Error(), sourceType)
		return ""
	}

	imgFallback := document.Find(imgFallbackSelector)

	for _, blacklistedSelector := range blacklistedSelectors {
		document.Find(blacklistedSelector).Each(func(i int, s *goquery.Selection) {
			removeNodes(s)
		})
	}

	foucsSelectorsUseful := len(focusSelectors) == 0
	for _, focusSelector := range focusSelectors {
		if focusSelector != "" {
			initialHTML, err := document.Find(focusSelector).Html()
			if err != nil {
				logger.Errorf("error while applying initial selector %s. error: %s", focusSelector, err.Error())
			}

			if len(strings.TrimSpace(initialHTML)) > 0 {
				document, err = goquery.NewDocumentFromReader(strings.NewReader(initialHTML))
				if err != nil {
					logger.Errorf("error while createing document from initial selector %s. error: %s", focusSelector, err.Error())
				}
				foucsSelectorsUseful = true
				break
			}
		}
	}

	// initital selector provided but non of them found any useful content
	if !foucsSelectorsUseful {
		return ""
	}

	document.Find("script,style,noscript").Each(func(i int, s *goquery.Selection) {
		removeNodes(s)
	})

	//output := bytes.NewBufferString("<div>")
	document.Find("*").Each(func(i int, s *goquery.Selection) {

		removeEmptyNodes(s)
		stripStyles(s)
		stripClasses(s)
		removeAdvertisementLeftovers(s)
		removeCommentNodes(s)
		removeAllDataDashAttrs(s)
	})

	//output.Write([]byte("</div>"))

	result, _ := document.Find("body").First().Html()
	//result := output.String()

	// add fallback image when no image is found in the article
	if !strings.Contains(result, "<img") && imgFallback != nil {
		if fallbackImgURL, exist := imgFallback.Attr("src"); exist {
			return "<div><img src='" + fallbackImgURL + "'/>" + result + "</div>"
		}
	}

	return result
}
