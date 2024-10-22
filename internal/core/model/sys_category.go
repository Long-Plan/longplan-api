package model

import "time"

type SysCategory struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	NameTH         string    `gorm:"column:name_th;type:varchar(255)" json:"name_th"`
	NameEN         string    `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CurriculumID   int       `gorm:"column:curriculum_id;type:int" json:"curriculum_id"`
	AtLeast        *bool     `gorm:"column:at_least;type:boolean" json:"at_least"`
	Credit         *int      `gorm:"column:credit;type:int" json:"credit"`
	TypeID         *int      `gorm:"column:category_type_id;type:int" json:"type_id"`
	IsDisplay      *bool     `gorm:"column:is_display;type:boolean" json:"is_display"`
	PrimaryColor   string    `gorm:"column:primary_color;type:varchar(20)" json:"primary_color"`
	SecondaryColor string    `gorm:"column:secondary_color;type:varchar(20)" json:"secondary_color"`
	Note           string    `gorm:"column:note;type:text" json:"note"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Type          SysCategoryType           `gorm:"foreignKey:TypeID;references:ID" json:"type"`
	Relationships []SysCategoryRelationship `gorm:"foreignKey:ParentCategoryID;references:ID" json:"relationships"`
	Requirements  []SysCategoryRequirement  `gorm:"foreignKey:CategoryID;references:ID" json:"requirements"`
	Courses       []SysCategoryCourse       `gorm:"foreignKey:CategoryID;references:ID" json:"courses"`
}
