package library

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	IsAvailable bool   `json:"available"`
}

type BookList struct {
	Books map[int]Book `json:"books"`
}

func NewBookList() *BookList {
	return &BookList{
		Books: make(map[int]Book),
	}
}

func AddBook(book *BookList, title, author string) {
	bookId := len(book.Books) + 1
	newBook := Book{ID: bookId, Title: title, Author: author, IsAvailable: false}
	book.Books[bookId] = newBook
	fmt.Println("Book added")
}

func RemoveBook(books *BookList, id int) {
	_, ok := books.Books[id]
	if !ok {
		fmt.Println("Book with ID", id, "not found")
		return
	}
	delete(books.Books, id)
	fmt.Println("Book removed")
}

func SearchBool(books *BookList, query string) []Book {
	var results []Book

	for _, book := range books.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
			results = append(results, book)
		}
	}
	return results
}

func CheckOutBook(books *BookList, id int) error {
	// Check if the book exists
	book, ok := books.Books[id]
	if !ok {
		return fmt.Errorf("Book with ID %d not found", id)
	}

	// Check if the book is available
	if !book.IsAvailable {
		return fmt.Errorf("Book with ID %d is already checked out", id)
	}

	// Update the availability status
	book.IsAvailable = false
	books.Books[id] = book
	fmt.Printf("Book '%s' checked out successfully\n", book.Title)
	return nil
}

func SaveBook(books *BookList) error {
	file, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("books.json", file, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Book saved")
	return nil

}
