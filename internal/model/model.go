package model

type Model struct {
	ID         uint32 `gorm:"primary_key"`
	CreatedBy  string
	ModifiedBy string
	CreatedOn  uint32
	ModifiedOn uint32
	IsDel      uint8
}
