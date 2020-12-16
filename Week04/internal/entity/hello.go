package entity

type Hello struct {
	ID   int64  `gorm:"column:id" json:"id" form:"id"`
	Name string `gorm:"column:name" json:"name" form:"name"`
}
