package models

import "time"

type Photo struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null; type:varchar(255)"`
	Caption   string    `gorm:"type:text" `
	PhotoUrl  string    `gorm:"not null; type:varchar(255)"`
	UserId    int       `gorm:"not null; type:int"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null, autoUpdateTime"`
}
