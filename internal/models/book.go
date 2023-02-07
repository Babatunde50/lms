package models

import (
	"context"
	"log"

	"time"
)

type Book struct {
	ID              int       `json:"id"`
	ISBN            int       `json:"isbn"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publication_date"`
	Publisher       string    `json:"publisher"`
	Genre           string    `json:"genre"`
	Language        string    `json:"language"`
	Description     string    `json:"description"`
	CoverImage      string    `json:"cover_image"`
	AvailableCopies int       `json:"available_copies"`
	TotalCopies     int       `json:"total_copies"`
	Format          string    `json:"format"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// add a book record to the database and returns an error if an error occur
func AddBook(data map[string]interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	b := Book{
		ISBN:            data["isbn"].(int),
		Title:           data["title"].(string),
		Author:          data["author"].(string),
		PublicationDate: data["publication_date"].(time.Time),
		Publisher:       data["publisher"].(string),
		Genre:           data["genre"].(string),
		Language:        data["language"].(string),
		Description:     data["description"].(string),
		CoverImage:      data["cover_image"].(string),
		AvailableCopies: data["available_copies"].(int),
		TotalCopies:     data["available_copies"].(int),
		Format:          data["format"].(string),
	}

	stmt := `INSERT INTO books 
			(isbn, title, author, publication_date, publisher, genre, language, description, cover_image, available_copies, total_copies,  format  )
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	`

	_, err := db.ExecContext(ctx, stmt, b.ISBN, b.Title, b.Author, b.PublicationDate, b.Publisher, b.Genre, b.Language, b.Description, b.CoverImage, b.AvailableCopies, b.TotalCopies, b.Format)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

// edit book
func UpdateBookAvailableCopies(remainingCopies, bookId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `UPDATE books 
			SET available_copies = $1
			WHERE id = $2
	`

	_, err := db.ExecContext(ctx, stmt, remainingCopies)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteBook(bookId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `DELETE FROM books 
			WHERE id = $1
		`

	_, err := db.ExecContext(ctx, stmt, bookId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GET book
func GetBook(bookId int) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var book Book

	stmt := `SELECT id, isbn, format, author, available_copies, cover_image, description, genre, language, total_copies, title, publisher, publication_date, created_at, updated_at FROM books 
			WHERE id = $1
		`

	row := db.QueryRowContext(ctx, stmt, bookId)

	err := row.Scan(&book.ID, &book.ISBN, &book.Format, &book.Author, &book.AvailableCopies, &book.CoverImage, &book.Description, &book.Genre, &book.Language, &book.TotalCopies, &book.Title, &book.Publisher, &book.PublicationDate, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		log.Println(err)
		return book, err
	}

	return book, nil
}

// get books
func GetBooks() ([]Book, error) {
	var books []Book

	return books, nil
}
