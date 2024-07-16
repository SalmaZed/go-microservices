package books_data

// book represents data about a book.
type Book struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Author            string  `json:"author"`
	YearOfPublication int     `json:"publication_year"`
	Price             float64 `json:"price"`
}

// books slice to seed record book data.
var Books = []Book{
	{ID: "1", Title: "Life Stories", Author: "Heather Newbold", YearOfPublication: 1998, Price: 4.99},
	{ID: "2", Title: "All Around The Town", Author: "Mary Higgins Clark", YearOfPublication: 2020, Price: 5.29},
	{ID: "3", Title: "The 48 Laws of Power", Author: "Robert Greene", YearOfPublication: 2023, Price: 10.99},
}
