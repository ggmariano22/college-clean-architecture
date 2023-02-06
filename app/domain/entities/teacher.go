package entities

type Teacher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
	Class string `json:"class"`
}
