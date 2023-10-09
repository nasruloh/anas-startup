package handler

import (
	"net/http"
	"startup-anas/helper"
	"startup-anas/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input dari user
	// map input dari user ke struct registrasiuserinput
	// struct di atas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil  {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response:= helper.APIResponse("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		
	}

	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil  {
		response:= helper.APIResponse("register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token, err := h,jvtService
	formatter := user.FormatUser(newUser, "tokentokentokentoken")
	response:= helper.APIResponse("account has been registed", http.StatusOK, "succes", formatter)
	c.JSON(http.StatusOK, response)
}