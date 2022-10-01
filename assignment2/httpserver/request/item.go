package request

import "assignment2/httpserver/repositories/models"

type CreateItemRequest struct {
	ItemCode    string `json:"itemcode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type UpdateItemRequest struct {
	ItemId      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (req *CreateItemRequest) ToModel(orderId int) (*models.Item, error) {
	item := &models.Item{
		OrderId:     orderId,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	return item, nil
}

func (req *UpdateItemRequest) ToModel(orderId int) (*models.Item, error) {
	item := &models.Item{
		ItemId:      req.ItemId,
		OrderId:     orderId,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	return item, nil
}
