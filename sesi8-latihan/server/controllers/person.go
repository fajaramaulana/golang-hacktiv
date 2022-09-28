package controllers

import (
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary Get all people
// @Decription Get list people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSuccess
// @Failure 404 {object} views.GetAllPeopleFailNotFound
// @Router /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusNotFound,
			Message: "GET_PEOPLE_FAIL",
			// Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// CreatePeople godoc
// @Summary Create new people
// @Decription Create new people
// @Tags person
// @Accept json
// @Produce json
// @Param people body models.Person true "People"
// @Success 200 {object} views.CreatePersonSuccess
// @Failure 400 {object} views.CreatePersonFail
// @Router /people [post]
func (p *PersonController) AddPeople(ctx *gin.Context) {
	var person models.Person

	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusBadRequest,
			Message: "ADD_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = p.db.Create(&person).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusBadRequest,
			Message: "ADD_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "ADD_PEOPLE_SUCCESS",
		Payload: person,
	})
}

// GetPersonById godoc
// @Summary Get person by id
// @Decription Get person by id
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} views.GetPersonByIdSuccess
// @Failure 404 {object} views.GetPersonByIdNotFound
// @Router /people/{id} [get]
func (p *PersonController) GetPersonById(ctx *gin.Context) {
	var person models.Person
	id := ctx.Param("id")

	err := p.db.First(&person, id).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusNotFound,
			Message: "GET_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PERSON_SUCCESS",
		Payload: person,
	})
}

// UpdatePerson godoc
// @Summary Update person by id
// @Decription Update person by id
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param person body models.Person true "Person"
// @Success 200 {object} views.UpdatePersonSuccess
// @Failure 400 {object} views.UpdatePersonFail
// @Router /people/{id} [put]
func (p *PersonController) UpdatePerson(ctx *gin.Context) {
	var person models.Person
	id := ctx.Param("id")

	err := p.db.First(&person, id).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusNotFound,
			Message: "UPDATE_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&person)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusBadRequest,
			Message: "UPDATE_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = p.db.Save(&person).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusBadRequest,
			Message: "UPDATE_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PERSON_SUCCESS",
		Payload: person,
	})
}

// DeletePerson godoc
// @Summary Delete person by id
// @Decription Delete person by id
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} views.DeletePersonSuccess
// @Failure 400 {object} views.DeletePersonFail
// @Router /people/{id} [delete]
func (p *PersonController) DeletePerson(ctx *gin.Context) {
	var person models.Person
	id := ctx.Param("id")

	err := p.db.First(&person, id).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusNotFound,
			Message: "DELETE_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	err = p.db.Delete(&person).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusBadRequest,
			Message: "DELETE_PERSON_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_PERSON_SUCCESS",
	})
}
