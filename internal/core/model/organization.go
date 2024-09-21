package model

type Organization struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	NameTH    string `gorm:"column:name_th;type:varchar(255)" json:"name_th"`
	NameEN    string `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"`
}