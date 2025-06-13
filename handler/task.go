package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/model"
	"todo-app/store"
)

type Handler struct {
	Store *store.MemoryStore
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	created := h.Store.Create(task)
	c.JSON(http.StatusOK, created)
}
func (h *Handler) ListTask(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.List())
}
func (h *Handler) DeleteTask(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ok := h.Store.Delete(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
func (h *Handler) MarkDone(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Store.MarkDone(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

//func (h *Handler) MarkDone(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPut {
//		http.Error(w, "Only PUT is allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	idStr := r.URL.Query().Get("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		http.Error(w, "Tasl not found", http.StatusNotFound)
//		return
//	}
//	ok := h.Store.MarkDone(id)
//	if ok != nil {
//		http.Error(w, "Task not found", http.StatusNotFound)
//		return
//	}
//	w.Header().Set("Content-Type", "appliscation/json")
//	json.NewEncoder(w).Encode(map[string]string{"status": "done"})
//}
