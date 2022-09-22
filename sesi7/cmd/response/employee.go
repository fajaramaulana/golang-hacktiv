package response

import "sesi7/cmd/entities"

type EmployeeFormatterCreate struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func FormatUserRegister(entities entities.Employee) EmployeeFormatterCreate {
	formatter := EmployeeFormatterCreate{
		ID:    entities.Id,
		Email: entities.Email,
	}

	return formatter
}
