package entities

import "time"

type Order struct {
	OrderId      int       `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
