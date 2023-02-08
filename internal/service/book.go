package service

import (
	"time"

	"github.com/Babatunde50/lms/internal/models"
)

type Book struct {
	ISBN            string
	Title           string
	Author          string
	PublicationDate time.Time
	Publisher       string
	Genre           string
	Language        string
	Description     string
	CoverImage      string
	AvailableCopies int
	TotalCopies     int
	Format          string
	Limit           int
	Offset          int
}

func (b *Book) Add() error {

	data := make(map[string]interface{})

	data["isbn"] = b.ISBN
	data["title"] = b.Title
	data["author"] = b.Author
	data["publication_date"] = b.PublicationDate
	data["publisher"] = b.Publisher
	data["genre"] = b.Genre
	data["language"] = b.Language
	data["description"] = b.Description
	data["cover_image"] = b.CoverImage
	data["available_copies"] = b.AvailableCopies
	data["total_copies"] = b.TotalCopies
	data["format"] = b.Format

	err := models.AddBook(data)
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) Edit() error {

	return nil
}

func (b *Book) Delete() error {

	return nil
}

func (b *Book) Get() (models.Book, error) {

	book, err := models.GetBook(b.ISBN)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *Book) GetAll() ([]models.Book, error) {

	books, err := models.GetBooks(b.Limit, b.Offset)
	if err != nil {
		return books, err
	}

	return books, nil
}
