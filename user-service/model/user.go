package model

type User struct {
	Uid         uint   `gorm:"primary_key"  json:"uid"`
	Username    string `gorm:"type:varchar(32);unique" json:"username"`
	Password    string `gorm:"type:varchar(64)" json:"_"`
	Mail        string `gorm:"type:varchar(200)" json:"mail"`
}