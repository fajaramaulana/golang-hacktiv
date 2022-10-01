package models

type Item struct {
	ItemId      int `gorm:"primaryKey;autoIncrement"`
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}
