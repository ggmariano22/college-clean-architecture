package entities

type Student struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         int    `json:"cpf"`
	PhoneNumber string `json:"phone_number"`
}

func NewStudent(name, email string, cpf int, phoneNumber string) *Student {
	return &Student{
		Name:        name,
		Email:       email,
		CPF:         cpf,
		PhoneNumber: phoneNumber,
	}
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetEmail() string {
	return s.Email
}

func (s *Student) GetCpf() int {
	return s.CPF
}

func (s *Student) GetPhoneNumber() string {
	return s.Name
}
