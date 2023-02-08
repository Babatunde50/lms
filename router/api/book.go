package api

import (
	"database/sql"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/Babatunde50/lms/internal/service"
	"github.com/Babatunde50/lms/pkg/app"
	setting "github.com/Babatunde50/lms/pkg/settings"
	"github.com/gin-gonic/gin"
)

type addBookRequest struct {
	ISBN            string                `json:"isbn" form:"isbn" binding:"required,isbn"`
	Title           string                `json:"title" form:"title" binding:"required"`
	Author          string                `json:"author" form:"author" binding:"required"`
	PublicationDate time.Time             `json:"publication_date" form:"publication_date" binding:"required" time_format:"2006-01-02" `
	Publisher       string                `json:"publisher" form:"publisher" binding:"required"`
	Genre           string                `json:"genre" form:"genre" binding:"required"`
	Language        string                `json:"language" form:"language" binding:"required"`
	Description     string                `json:"description" form:"description" binding:"required"`
	AvailableCopies int                   `json:"available_copies" form:"available_copies" binding:"required"`
	TotalCopies     int                   `json:"total_copies" form:"total_copies" binding:"required,gtefield=AvailableCopies"`
	Format          string                `json:"format" form:"format" binding:"required,oneof=hard_copy soft_copy"`
	CoverImage      *multipart.FileHeader `form:"cover_image" json:"cover_image" binding:"required"`
}

func AddBook(c *gin.Context) {
	appG := app.Gin{C: c}
	var req addBookRequest
	err := appG.C.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}
	dst := setting.AppSetting.ImageSavePath + req.CoverImage.Filename
	err = appG.C.SaveUploadedFile(req.CoverImage, dst)

	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	bookService := service.Book{
		ISBN:            req.ISBN,
		Author:          req.Author,
		AvailableCopies: req.AvailableCopies,
		Title:           req.Title,
		PublicationDate: req.PublicationDate,
		Publisher:       req.Publisher,
		Genre:           req.Genre,
		Language:        req.Language,
		TotalCopies:     req.TotalCopies,
		Format:          req.Format,
		CoverImage:      dst,
		Description:     req.Description,
	}

	err = bookService.Add()
	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	appG.Response(http.StatusCreated, http.StatusCreated, "success", nil)

}

type getBookRequest struct {
	ISBN string `uri:"isbn" binding:"required,isbn"`
}

func GetBook(c *gin.Context) {
	appG := app.Gin{C: c}

	var req getBookRequest

	err := appG.C.ShouldBindUri(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	bookService := service.Book{ISBN: req.ISBN}

	book, err := bookService.Get()

	if err != nil {

		if err == sql.ErrNoRows {
			appG.Response(http.StatusNotFound, http.StatusNotFound, err.Error(), nil)
			return
		}

		appG.Response(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, http.StatusOK, "success", book)

}

type getBooksRequest struct {
	PageId   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
}

func GetBooks(c *gin.Context) {
	appG := app.Gin{C: c}

	var req getBooksRequest

	err := appG.C.ShouldBindQuery(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil)
		return
	}

	bookService := service.Book{Limit: req.PageSize, Offset: (req.PageId - 1) * req.PageSize}

	books, err := bookService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, http.StatusOK, "success", books)

}
