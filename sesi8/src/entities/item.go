package entities

import "time"

type Item struct {
	ItemId      int       `json:"item_id" gorm:"primaryKey;autoIncrement"`
	ItemCode    string    `json:"item_code" gorm:"unique"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
