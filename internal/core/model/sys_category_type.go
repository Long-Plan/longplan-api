package model

import "time"

type SysCategoryType struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	NameTH    string    `gorm:"column:name_th;type:varchar(255);unique" json:"name_th"`
	NameEN    string    `gorm:"column:name_en;type:varchar(255);unique" json:"name_en"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
