package request

import (
	"assignment2/httpserver/repositories/models"
	"time"
)

type CreateOrderRequest struct {
	CustomerName string              `json:"customerName" example:"John Doe"`
	OrderAt      string              `json:"orderedAt" example:"2021-05-01T00:00:00Z"`
	Items        []CreateItemRequest `json:"items"`
}

type UpdateOrderRequest struct {
	// OrderID      int                 `json:"orderid"`
	CustomerName string              `json:"customerName" example:"Fajar"`
	OrderAt      string              `json:"orderedAt" example:"2022-10-01T21:21:21Z"`
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
