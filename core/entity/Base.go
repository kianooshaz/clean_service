package entity

import (
	"gorm.io/gorm"
	"time"
)

// midonam inja clean nist , vali in yeki ja ghabol bokon :), namikhastam ye struct baraie gorm dorost bokonam :(
// todo hala badan dorostesh mikonam
type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
