package valueObjects

type Person struct {
	name  string
	email string
	cpf   int
}

func NewPerson(name, email string, cpf int) *Person {
	return &Person{
		name:  name,
		email: email,
		cpf:   cpf,
	}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetEmail() string {
	return p.email
}

func (p *Person) GetCpf() int {
	return p.cpf
}
