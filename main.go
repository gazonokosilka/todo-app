package main

import (
	"log"
	"todo-app/handler"
	"todo-app/store"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db, err := store.NewPostgresStore("host=localhost port=5432 user=postgres password=admin dbname=todo sslmode=disable")

	if err != nil {
		log.Fatal("Ошибка подключение к БД", err)
	}
	taskStore := db

	h := &handler.Handler{Store: taskStore}

	r := gin.Default()

	r.POST("/create", h.CreateTask)
	r.GET("/list", h.ListTask)
	r.DELETE("/delete", h.DeleteTask)
	r.PUT("/done", h.MarkDone)

	r.Run(":8080")
}
