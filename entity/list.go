package entity

type ArticleListItem struct {
	Title       string
	Description string
	Children    []ArticleListItem
}
