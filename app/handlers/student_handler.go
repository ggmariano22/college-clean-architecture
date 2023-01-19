package handlers

import (
	"college/app/domain/entities"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type StudentUseCase interface {
	AddStudent(student *entities.Student)
	GetStudents() []*entities.Student
}

type StudentHandler struct {
	useCase StudentUseCase
}

func NewStudentHandler(useCase StudentUseCase) *StudentHandler {
	return &StudentHandler{
		useCase: useCase,
	}
}

func (h *StudentHandler) AddStudentHandler(c *gin.Context) {
	student := &entities.Student{}

	if err := json.NewDecoder(c.Request.Body).Decode(student); err != nil {
		fmt.Println("Houve um erro ao processar o body da request: ", err)
		os.Exit(1)
	}

	h.useCase.AddStudent(student)

	c.JSON(201, gin.H{
		"message": fmt.Sprintf("Aluno %s criado com sucesso!", student.Name),
	})
}

func (h *StudentHandler) GetStudentsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": h.useCase.GetStudents(),
	})
}
