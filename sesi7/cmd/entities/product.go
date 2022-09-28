package entities

type Task struct {
	Id           int    `json:"id"`
	Product_name string `json:"productName"`
	Price        int    `json:"price"`
	Description  string `json:"description"`
	Division     string `json:"division"`
}
