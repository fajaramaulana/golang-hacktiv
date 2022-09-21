package entities

type Employee struct {
	Id        int    `json:"id"`
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Division  string `json:"division"`
}
