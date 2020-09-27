package model

type Article struct {
	*Model
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         uint8
}
