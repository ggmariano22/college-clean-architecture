package handlers

import (
	"college/app/handlers/dtos"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentUseCase interface {
	AddStudent(student *dtos.StudentInput)
	GetStudents() []*dtos.StudentOutput
	GetStudentByCPF(cpf int) (*dtos.StudentOutput, error)
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
	student := &dtos.StudentInput{}

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
	fmt.Println(h.useCase.GetStudents())
	c.JSON(200, gin.H{
		"data": h.useCase.GetStudents(),
	})
}

func (h *StudentHandler) GetStudentByCPF(c *gin.Context) {
	cpf, err := strconv.Atoi(c.Param("cpf"))

	if err != nil {
		c.JSON(422, gin.H{
			"data": fmt.Sprint("Valor para parâmetro CPF inválido."),
		})
	}

	student, err := h.useCase.GetStudentByCPF(cpf)

	if err != nil {

		c.JSON(404, gin.H{
			"data": fmt.Sprint(err),
		})

		return
	}

	c.JSON(200, gin.H{
		"data": student,
	})
}
