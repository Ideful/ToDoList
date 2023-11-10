package handler

import (
	model "ToDoList/internal/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h* Handler) signUp(c *gin.Context) {
	var input model.User
	fmt.Print(2)
	if err := c.BindJSON(&input); err != nil {
		log.Fatal(err)
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err:=h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id":id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (h* Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		log.Fatal(err)
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	token,err:=h.services.Authorization.GenerateToken(input.Username,input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"token":token,
	})
}