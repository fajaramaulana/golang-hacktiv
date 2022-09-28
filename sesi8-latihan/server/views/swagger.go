package views

import "sesi8-latihan/server/models"

type GetAllPeopleSuccess struct {
	Message string      `json:"message" example:"GET_PEOPLE_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
	Payload models.Person
}

type GetAllPeopleFailNotFound struct {
	StatusCode int    `json:"status" example:"404"`
	Message    string `json:"message" example:"GET_PEOPLE_FAIL"`
}

type CreatePersonSuccess struct {
	Message string      `json:"message" example:"CREATE_PERSON_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
	Payload models.Person
}

type CreatePersonFail struct {
	StatusCode int    `json:"status" example:"400"`
	Message    string `json:"message" example:"ADD_PEOPLE_FAIL"`
}

type GetPersonByIdSuccess struct {
	Message string      `json:"message" example:"GET_PERSON_BY_ID_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
	Payload models.Person
}

type GetPersonByIdNotFound struct {
	StatusCode int    `json:"status" example:"404"`
	Message    string `json:"message" example:"GET_PERSON_FAIL"`
	Error      string `json:"error" example:"record not found"`
}

type UpdatePersonFail struct {
	StatusCode int    `json:"status" example:"400"`
	Message    string `json:"message" example:"UPDATE_PERSON_FAIL"`
	Error      string `json:"error"`
}

type UpdatePersonSuccess struct {
	Message string      `json:"message" example:"UPDATE_PERSON_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
	Payload models.Person
}

type DeletePersonSuccess struct {
	Message string      `json:"message" example:"DELETE_PERSON_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
}

type DeletePersonFail struct {
	StatusCode int    `json:"status" example:"400"`
	Message    string `json:"message" example:"DELETE_PERSON_FAIL"`
	Error      string `json:"error"`
}
