package models

type ArticleSource struct {
	ID                  int
	Source              string
	SourceImageURL      string
	SourceDisplayName   string
	AvailableCategories []int
}

type ArticleCategory struct {
	ID                  int
	Category            string
	CategoryDisplayName string
	IsPublic            bool
}

type AppSettings struct {

	// ArticleCategory contains all the article categories
	ArticleCategory []ArticleCategory

	// ArticleSource contains all the article sources
	ArticleSource []ArticleSource
}
