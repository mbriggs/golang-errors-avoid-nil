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

// ErrBookNotFound is returned when a book is not found in the library
var ErrBookNotFound = errors.New("book not found in library")

// ErrInvalidBook is returned when trying to add a book without a title
var ErrInvalidBook = errors.New("invalid book")

// ErrBookAlreadyExists is returned when trying to add a book that already exists
var ErrBookAlreadyExists = errors.New("book already exists")

// Library is a collection of books
type Library struct {
	Books []Book
}

// AddBook to Library
func (l *Library) AddBook(book Book) error {
	if book.Title == "" {
		return ErrInvalidBook
	}

	for _, b := range l.Books {
		if b.Title == book.Title {
			return fmt.Errorf("adding %s: %w", book.Title, ErrBookAlreadyExists)
		}
	}

	l.Books = append(l.Books, book)

	return nil
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

	return Book{}, fmt.Errorf("finding %s: %w", title, ErrBookNotFound)
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
	library := Library{}

	// books to add
	books := []Book{
		{Title: "1984"},
		{Title: "Dune"},
		{Title: "Hitchhiker's Guide to the Galaxy"},
		{Title: "The Lord of the Rings"},
	}

	for _, book := range books {
		err := library.AddBook(book)

		// if a book already exists, log out and continue
		if errors.Is(err, ErrBookAlreadyExists) {
			fmt.Printf("Book already exists! %s\n", err)
			continue
		}

		// if book is invalid, log out and exit
		if err != nil {
			fmt.Printf("Invalid book! %s\n", err)
			os.Exit(1)
			return
		}
	}

	searchTitles := os.Args[1:]

	for _, searchTitle := range searchTitles {
		// Try to find a book
		book, err := library.FindBook(searchTitle)

		// Book not found should log out and exit
		if errors.Is(err, ErrBookNotFound) {
			fmt.Printf("Book not found! %s\n", err)
			os.Exit(1)
			return
		}

		// Other reasons should be a crash
		if err != nil {
			panic(err)
		}

		// Success condition
		fmt.Printf("Found %s!\n", book)
	}
}
