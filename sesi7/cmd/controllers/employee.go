package controllers

import (
	"fmt"
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

func NewEmployeeController(employeeService services.EmployeeService) *employeeController {
	return &employeeController{employeeService}
}

func (employeeController *employeeController) EmployeeRoutes(group *gin.RouterGroup) {
	route := group.Group("/employee")
	route.POST("/create", employeeController.AddEmployee)
	route.GET("/getall", employeeController.GetAllEmployee)
}

func (h *employeeController) AddEmployee(c *gin.Context) {
	var input request.RegisterEmployee

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%#v \n", input)

	newEmployee, err := h.employeeService.CreateEmployee(input)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": newEmployee})
}

func (h *employeeController) GetAllEmployee(c *gin.Context) {
	employees, err := h.employeeService.GetAllEmployee()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": employees})
}
