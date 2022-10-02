package request

import "assignment2/httpserver/repositories/models"

type CreateItemRequest struct {
	ItemCode    string `json:"itemcode" example:"A001"`
	Description string `json:"description" example:"Baju"`
	Quantity    int    `json:"quantity" example:"1"`
}

type UpdateItemRequest struct {
	ItemId      int    `json:"itemId" example:"1"`
	ItemCode    string `json:"itemCode" example:"A001"`
	Description string `json:"description" example:"Baju"`
	Quantity    int    `json:"quantity" example:"1"`
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
