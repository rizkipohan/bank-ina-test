package delivery

import (
	"bank-ina-test/domain"
	"bank-ina-test/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usecase usecase.UserUsecase
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Usecase.RegisterUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Register successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Login(c *gin.Context) {
	var user domain.Login
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.Usecase.LoginUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respData := make(map[string]string)
	respData["token"] = tokens

	var resp domain.ResponseModels
	resp.Message = "Login successful"
	resp.Data = respData

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Usecase.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "OK"
	resp.Data = users

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.Usecase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "OK"
	resp.Data = user

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user.ID = uint(id)

	err = h.Usecase.UpdateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Update successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.Usecase.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp domain.ResponseModels
	resp.Message = "Delete successful"
	resp.Data = nil

	c.JSON(http.StatusOK, resp)
}
