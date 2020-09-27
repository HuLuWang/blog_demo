package model

type Tag struct {
	*Model
	Name  string
	State uint8
}
