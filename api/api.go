package api

import (
	"log"

	"75.GO/api/handler"
	"75.GO/internal/mongodb"
	"github.com/gin-gonic/gin"
)

func Routes(db *mongodb.StudentMongoDb) {
	router := gin.Default()
	studentHandler := handler.NewUserHandler(db)
	router.POST("/students", studentHandler.CreateStudent)
	router.GET("/students/:id", studentHandler.GetByIdStudent)
	router.PUT("/students/:id", studentHandler.UpdateStudent)
	router.DELETE("/students/:id", studentHandler.DeleteStudent)
	log.Fatal(router.Run(":8888"))

}
