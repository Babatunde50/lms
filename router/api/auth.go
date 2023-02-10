package api

import (
	"database/sql"
	"net/http"

	"github.com/Babatunde50/lms/internal/service"
	"github.com/Babatunde50/lms/pkg/app"
	setting "github.com/Babatunde50/lms/pkg/settings"
	"github.com/Babatunde50/lms/pkg/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

type loginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	var data loginForm
	if err := c.ShouldBindJSON(&data); err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check if email is in database...
	memberService := service.Member{Email: data.Email}

	member, err := memberService.FindByEmail()

	if err != nil {
		if err == sql.ErrNoRows {
			appG.Response(http.StatusNotFound, http.StatusNotFound, err.Error(), nil)
			return
		}
		appG.Response(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	// check password is correct
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(data.Password))

	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	token, err := token.JWTTokenMaker.CreateToken(data.Email, setting.TokenSetting.TokenExpireDuration)

	if err != nil {
		appG.Response(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, http.StatusAccepted, "Sign up success", map[string]string{"access_token": token})

}
