package controllers

import (
	"assignment2/httpserver/controllers/views"
	"assignment2/httpserver/request"
	"assignment2/httpserver/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req request.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, views.M_BAD_REQUEST)
		return
	}

	response := c.service.CreateOrder(&req)

	WriteJsonResponse(ctx, response)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var req request.UpdateOrderRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, views.M_BAD_REQUEST)
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, views.M_BAD_REQUEST)
		return
	}

	response := c.service.UpdateOrder(&req, id)

	WriteJsonResponse(ctx, response)
}

func (c *OrderController) GetAll(ctx *gin.Context) {
	response := c.service.FindAllOrder()
	WriteJsonResponse(ctx, response)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, views.M_BAD_REQUEST)
		return
	}

	response := c.service.DeleteOrder(id)

	WriteJsonResponse(ctx, response)
}
