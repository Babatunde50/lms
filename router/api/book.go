package api

import (
	"net/http"
	"strconv"

	"github.com/Babatunde50/lms/internal/service"
	"github.com/Babatunde50/lms/pkg/app"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	appG := app.Gin{C: c}

	bookId, err := strconv.Atoi(appG.C.Param("id"))

	if err != nil {
		appG.Response(400, 400, "can't convert to int", nil)
		return
	}

	valid := validation.Validation{}
	valid.Min(bookId, 1, "id")

	if valid.HasErrors() {
		// implement error logger...
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, valid.Errors[0].Message, nil)
		return
	}

	bookService := service.Book{ID: bookId}

	book, err := bookService.Get()

	if err != nil {
		appG.Response(400, 400, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, http.StatusOK, "success", book)

}
