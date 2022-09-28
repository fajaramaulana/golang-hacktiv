package controllers

import (
	"fmt"
	"net/http"
	"sesi7/cmd/repositories"
	"sesi7/cmd/request"
	"sesi7/cmd/response"
	"sesi7/cmd/services"
	"sesi7/configs/utils"

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

	dataResponse := response.FormatUserRegister(newEmployee)

	response := utils.APIResponse("Data Berhasil ditambahkan", http.StatusOK, "Success", dataResponse)

	c.JSON(http.StatusOK, response)

}

// GetEmployee godoc
// @Summary Get all employee
// @Description Get all employee
// @Tags Employee
// @Accept  json
// @Produce  json
// @Success 200 {object} Employee
// @Router /employee/getall [get]
func (h *employeeController) GetAllEmployee(c *gin.Context) {
	employees, err := h.employeeService.GetAllEmployee()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := utils.APIResponse("Data Berhasil Ditemukan", http.StatusOK, "Success", employees)

	c.JSON(http.StatusOK, response)
}
