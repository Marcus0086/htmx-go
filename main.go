package main

import (
	"htmx-go/internal/app/handlers"
	"htmx-go/internal/app/repositories"
	"htmx-go/internal/app/services"
	"htmx-go/internal/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbPool.Close()
	todoRepo := repositories.NewTodoRepository(dbPool)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", todoHandler.ListTodosHandler)
	router.POST("/add", todoHandler.CreateTodoHandler)
	router.POST("/update", todoHandler.MarkTodoDoneHandler)
	router.POST("/delete", todoHandler.DeleteTodoHandler)

	log.Println("Starting server on :3000...")
	log.Fatal(router.Run(":3000"))
}
