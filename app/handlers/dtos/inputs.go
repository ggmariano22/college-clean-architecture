package dtos

type CommonPersonInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

type StudentInput struct {
	CommonPersonInput
	PhoneNumber string `json:"phone_number"`
}

type TeacherInput struct {
	CommonPersonInput
	Class string `json:"class"`
}
