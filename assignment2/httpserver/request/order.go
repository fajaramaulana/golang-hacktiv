package request

import (
	"assignment2/httpserver/repositories/models"
	"time"
)

type CreateOrderRequest struct {
	CustomerName string              `json:"customerName"`
	OrderAt      string              `json:"orderedAt"`
	Items        []CreateItemRequest `json:"items"`
}

type UpdateOrderRequest struct {
	OrderID      int                 `json:"orderid"`
	CustomerName string              `json:"customerName"`
	OrderAt      string              `json:"orderedAt"`
	Items        []UpdateItemRequest `json:"items"`
}

func (req *CreateOrderRequest) ToModel() (*models.Order, error) {
	convertTime, err := time.Parse(time.RFC3339, req.OrderAt)
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    convertTime,
	}

	return order, nil
}

func (req *UpdateOrderRequest) ToModel(id int) (*models.Order, error) {
	convertTime, err := time.Parse(time.RFC3339, req.OrderAt)
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		OrderId:      id,
		CustomerName: req.CustomerName,
		OrderedAt:    convertTime,
	}

	return order, nil
}
