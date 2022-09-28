package models

type Person struct {
	// gorm.Model
	Id      int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
