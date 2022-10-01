package gorm

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{db}
}

func (r *orderRepo) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) Update(order *models.Order) error {
	return r.db.First(&models.Order{}, order.OrderId).Updates(order).Error
}

func (r *orderRepo) Delete(id int) error {
	return r.db.Delete(&models.Order{}, id).Error
}

func (r *orderRepo) FindAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	return orders, err
}
