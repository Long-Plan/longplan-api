package model

import "time"

type SysCategoryRequirement struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int       `gorm:"column:category_id;type:int;not null" json:"category_id"`
	Regex      *string   `gorm:"column:regex;type:varchar(255)" json:"regex,omitempty"`
	Credit     *int      `gorm:"column:credit;type:int" json:"credit,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
