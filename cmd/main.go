package main

import (
	"college/app/domain/repositories"
	"college/app/domain/useCases"
	"college/app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	studentRepository := repositories.NewStudentRepository()
	studentUseCase := useCases.NewStudentUseCase(studentRepository)
	studentHandler := handlers.NewStudentHandler(studentUseCase)

	r := gin.Default()

	r.POST("/student", studentHandler.AddStudentHandler)
	r.GET("/students", studentHandler.GetStudentsHandler)

	r.Run(":8995")
}
