package model

type ArticleTag struct {
	*Model
	ArticleID uint32
	TagID     uint32
}
