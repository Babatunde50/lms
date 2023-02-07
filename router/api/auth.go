package api

import (
	"net/http"

	"github.com/Babatunde50/lms/internal/service"
	"github.com/Babatunde50/lms/pkg/app"
	"github.com/gin-gonic/gin"
)

type signUpForm struct {
	Name        string `form:"name" json:"name" binding:"required,min=3"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	Password    string `form:"password" json:"password" binding:"required,min=6"`
	Address     string `form:"address" json:"address" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required,min=10"`
}

func SignUp(c *gin.Context) {
	appG := app.Gin{C: c}

	var data signUpForm
	if err := c.ShouldBindJSON(&data); err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	memberService := service.Member{
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
	}

	err := memberService.Add()

	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, http.StatusAccepted, "Sign up success", nil)

}

func Login(c *gin.Context) {

}
