package main

import (
	"fmt"
	"os"

	"example.com/book/library"
)

func main() {
	// bookList := &library.BookList{}
	bookList := library.NewBookList()

	for {
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Search Books")
		fmt.Println("4. Check Out Book")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title, author string
			fmt.Print("Enter title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter author: ")
			fmt.Scanln(&author)
			library.AddBook(bookList, title, author)
		case 2:
			var id int
			fmt.Scanln(&id)
			library.RemoveBook(bookList, id)
		case 3:
			var search string
			fmt.Println("Search for book title: ")
			fmt.Scanln(&search)
			library.SearchBool(bookList, search)
		case 4:
			var id int
			fmt.Scanln(&id)
			library.CheckOutBook(bookList, id)
		case 5:
			fmt.Println("Exiting...")
			library.SaveBook(bookList)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
