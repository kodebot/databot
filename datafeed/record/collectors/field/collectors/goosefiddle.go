package collectors

import (
	goose "github.com/advancedlogic/GoOse"
)

func ExtractContentGoose(url string) (*goose.Article, error) {
	g := goose.New()
	return g.ExtractFromURL(url)
	// println("title", article.Title)
	// println("description", article.MetaDescription)
	// println("keywords", article.MetaKeywords)
	// println("content", article.CleanedText)
	// println("url", article.FinalURL)
	// println("top image", article.TopImage)
}
