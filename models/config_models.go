package models

// ArticleSource defines the structure for article source
type ArticleSource struct {
	ID                  int
	Source              string
	SourceImageURL      string
	SourceDisplayName   string
	AvailableCategories []int
}

// ArticleCategory defines the structure for article category
type ArticleCategory struct {
	ID                  int
	Category            string
	CategoryDisplayName string
	IsPublic            bool
}

// AppSettings provide application wide configurable settings
type AppSettings struct {
	ConnectionString string
	LoadArticlesCron string

	// ArticleCategory contains all the article categories
	ArticleCategory []ArticleCategory

	// ArticleSource contains all the article sources
	ArticleSource []ArticleSource
}
