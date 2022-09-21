package controllers

import (
	"sesi7/cmd/repositories"
	"sesi7/cmd/request"
	"sesi7/cmd/services"

	"github.com/gin-gonic/gin"
)

type employeeController struct {
	employeeService services.EmployeeService
}

type service struct {
	repository repositories.RepositoryEmployee
}

type EmployeeService interface {
}

func NewEmployeeController(employeeService services.EmployeeService) *employeeController {
	return &employeeController{}
}

func (employeeController *employeeController) EmployeeRoutes(group *gin.RouterGroup) {
	group.POST("/employee", employeeController.AddEmployee)
}

func (h *employeeController) AddEmployee(c *gin.Context) {
	var input request.RegisterEmployee

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newEmployee, err := h.employeeService.CreateEmployee(input)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": newEmployee})
}
