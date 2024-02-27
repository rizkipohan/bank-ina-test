package delivery

import (
	"bank-ina-test/domain"
	"bank-ina-test/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Usecase usecase.TaskUsecase
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Usecase.CreateTask(c.Request.Context(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Create successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.Usecase.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "OK"
	resp.Data = tasks

	c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.Usecase.GetTaskByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "OK"
	resp.Data = task

	c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task.ID = uint(id)

	err = h.Usecase.UpdateTask(c.Request.Context(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Update successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = h.Usecase.DeleteTask(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Delete successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}
