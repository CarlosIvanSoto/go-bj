package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model

	Nickname string `gorm:"index:idx_name,unique" json:"nickname"`
	Credits  uint64 `json:"credits"`
}
