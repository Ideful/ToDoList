package handler

import (
	model "ToDoList/internal/models"
	"log"
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
)

func (h* Handler) signUp(c *gin.Context) {
	var input model.User
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

func (h* Handler) signIn(c *gin.Context) {
	
}