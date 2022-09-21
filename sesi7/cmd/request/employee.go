package request

type RegisterEmployee struct {
	Full_name string `json:"fullname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Division  string `json:"division"`
}
