package gofakeit

import "math/rand"

func BookTitle() string { return bookTitle(globalFaker.Rand) }

func (f *Faker) BookTitle() string { return bookTitle(f.Rand) }

func bookTitle(r *rand.Rand) string { return getRandValue(r, []string{"book", "title"}) }

func BookAuthor() string { return bookAuthor(globalFaker.Rand) }

func (f *Faker) BookAuthor() string { return bookAuthor(f.Rand) }

func bookAuthor(r *rand.Rand) string { return getRandValue(r, []string{"book", "author"}) }

func BookGenre() string { return bookGenre(globalFaker.Rand) }

func (f *Faker) BookGenre() string { return bookGenre(f.Rand) }

func bookGenre(r *rand.Rand) string { return getRandValue(r, []string{"book", "genre"}) }

type BookInfo struct {
	Title  string `json:"title" xml:"name"`
	Author string `json:"author" xml:"author"`
	Genre  string `json:"genre" xml:"genre"`
}

func Book() *BookInfo { return book(globalFaker.Rand) }

func (f *Faker) Book() *BookInfo { return book(f.Rand) }

func book(r *rand.Rand) *BookInfo {
	return &BookInfo{
		Title:  bookTitle(r),
		Author: bookAuthor(r),
		Genre:  bookGenre(r),
	}
}

func addBookLookup() {
	AddFuncLookup("book", Info{
		Display:     "Book",
		Category:    "book",
		Description: "Generates a comprehensive book dataset, including title, author, and genre",
		Example: `{
	"title": "Anna Karenina",
	"author": "Toni Morrison",
	"genre": "Thriller"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return book(r), nil
		},
	})

	AddFuncLookup("booktitle", Info{
		Display:     "Title",
		Category:    "book",
		Description: "Generates a book title from a diverse range of genres and periods",
		Example:     "Hamlet",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return bookTitle(r), nil
		},
	})

	AddFuncLookup("bookauthor", Info{
		Display:     "Author",
		Category:    "book",
		Description: "Generates the name of a book author from various times and places",
		Example:     "Mark Twain",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return bookAuthor(r), nil
		},
	})

	AddFuncLookup("bookgenre", Info{
		Display:     "Genre",
		Category:    "book",
		Description: "Generates a genre from a wide range of book categories",
		Example:     "Adventure",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return bookGenre(r), nil
		},
	})
}
