package services

import (
	"assignment2/httpserver/controllers/views"
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"assignment2/httpserver/request"
	"net/http"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
	itemRepo  repositories.ItemRepo
}

func NewOrderService(orderRepo repositories.OrderRepo, itemRepo repositories.ItemRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo, itemRepo: itemRepo}
}

func (service *OrderService) CreateOrder(req *request.CreateOrderRequest) *views.Response {
	order, err := req.ToModel()
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
	}

	err = service.orderRepo.Create(order)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	items := []models.Item{}

	for _, v := range req.Items {
		itemModel, err := v.ToModel(order.OrderId)
		if err != nil {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		items = append(items, *itemModel)
	}

	err = service.itemRepo.CreateManyItem(items)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, order)
}

func (service *OrderService) UpdateOrder(req *request.UpdateOrderRequest, id int) *views.Response {
	order, err := req.ToModel(id)
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
	}

	err = service.orderRepo.Update(order)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	for _, v := range req.Items {
		itemModel, err := v.ToModel(order.OrderId)
		if err != nil {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}

		err = service.itemRepo.Update(itemModel)

		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, order)
}

func (service *OrderService) FindAllOrder() *views.Response {
	orders, err := service.orderRepo.FindAll()

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	listOrders := make([]views.OrderGetAll, 0)

	for _, v := range orders {
		items, err := service.itemRepo.FindByOrderId(v.OrderId)
		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}

		listItem := make([]views.ItemGetAll, 0)

		for _, x := range items {
			listItem = append(listItem, views.ItemGetAll{
				ItemId:      x.ItemId,
				ItemCode:    x.ItemCode,
				Description: x.Description,
				Quantity:    x.Quantity,
			})
		}

		listOrder := views.OrderGetAll{
			OrderId:      v.OrderId,
			OrderedAt:    v.OrderedAt,
			CustomerName: v.CustomerName,
			Items:        listItem,
		}

		listOrders = append(listOrders, listOrder)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, listOrders)
}

func (service *OrderService) DeleteOrder(id int) *views.Response {
	err := service.orderRepo.Delete(id)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = service.itemRepo.DeleteItemByOrderId(id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, nil)
}
