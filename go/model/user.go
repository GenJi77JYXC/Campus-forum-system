package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"size:512"` // 哈希加密
	State    bool
}
