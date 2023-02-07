package service

import "github.com/Babatunde50/lms/internal/models"

type Book struct {
	ID int
}

func (b *Book) Add() error {

	return nil
}

func (b *Book) Edit() error {

	return nil
}

func (b *Book) Delete() error {

	return nil
}

func (b *Book) Get() (models.Book, error) {

	book, err := models.GetBook(b.ID)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *Book) GetAll() ([]models.Book, error) {

	var books []models.Book

	return books, nil
}
