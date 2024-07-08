package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Book is a book in a library
type Book struct {
	Title string
}

func (b Book) String() string {
	return b.Title
}

// ErrNotFound is returned when a book is not found in the library
var ErrNotFound = errors.New("book not found in library")

// Library is a collection of books
type Library struct {
	Books []Book
}

// FindBook in Library by title. If not found, returns ErrNotFound
func (l Library) FindBook(title string) (Book, error) {
	searchTitle := normalizeTitle(title)

	for _, book := range l.Books {
		bookTitle := normalizeTitle(book.Title)

		if strings.Contains(bookTitle, searchTitle) {
			// Found the book
			return book, nil
		}
	}

	return Book{}, fmt.Errorf("finding %s: %w", title, ErrNotFound)
}

// normalizeTitle normalizes a title by making it lowercase and removing punctuation and whitespace
func normalizeTitle(title string) string {
	var searchTitle strings.Builder

	for _, r := range title {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			continue
		}
		searchTitle.WriteRune(unicode.ToLower(r))
	}

	return searchTitle.String()
}

func main() {

	// Create a Library with some books
	library := Library{
		Books: []Book{
			{Title: "1984"},
			{Title: "Dune"},
			{Title: "Hitchhiker's Guide to the Galaxy"},
			{Title: "The Lord of the Rings"},
		},
	}

	searchTitles := os.Args[1:]

	for _, searchTitle := range searchTitles {
		// Try to find a book
		book, err := library.FindBook(searchTitle)

		// Book not found should log out and exit
		if errors.Is(err, ErrNotFound) {
			fmt.Printf("Book not found! %s\n", err)
			os.Exit(1)
		}

		// Other reasons should be a crash
		if err != nil {
			panic(err)
		}

		// Success condition
		fmt.Printf("Found %s!\n", book)
	}
}
