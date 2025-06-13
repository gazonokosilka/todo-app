package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/handler"
	"todo-app/store"
)

func main() {
	r := gin.Default()
	memoryStore := store.NewMemoryStore()
	h := handler.Handler{Store: memoryStore}

	r.POST("/create", h.CreateTask)
	r.GET("/list", h.ListTask)
	r.PUT("/done", h.MarkDone)
	r.DELETE("/delete", h.DeleteTask)

	r.Run(":8080")
}
