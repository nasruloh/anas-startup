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

func (h *userHandler) Login(c *gin.Context){
	//user memasukkan input  (email & password)
	//input di tangkap handler
	//mapping dari user ke input struct
	//input struct passing service
	//di service mencari dg bantuan repository user dengan email x
	//mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response:= helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response:= helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")
	response:= helper.APIResponse("Succesfuly loggedin", http.StatusOK, "succes", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailablity(c *gin.Context){
	// ada input email dari user
	// input email di ampping ke struct input
	// strcut input di-pasing ke service
	// service akan memanggil repostory - email sudah ada atau belum
	//repository - db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response:= helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "server error"}
		response:= helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	metaMessage := "email has been registered" 

	if IsEmailAvailable {
		metaMessage = "email is available"
	}  

	response:= helper.APIResponse(metaMessage, http.StatusUnprocessableEntity, "error", data)
	c.JSON(http.StatusUnprocessableEntity, response)
	return

}