package model

type Account struct {
	CMUITAccount string `gorm:"primaryKey;column:cmuitaccount;type:varchar(255)" json:"cmuitaccount"`
	Prename      string `gorm:"type:varchar(255)" json:"prename"`
	Firstname    string `gorm:"type:varchar(255);not null" json:"firstname"`
	Lastname     string `gorm:"type:varchar(255);not null" json:"lastname"`
	AccountType  string `gorm:"type:varchar(255);not null" json:"account_type"`
	Organization string `gorm:"type:varchar(255);not null" json:"organization"`
	CreatedAt    string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    string `gorm:"autoUpdateTime" json:"updated_at"`
}
