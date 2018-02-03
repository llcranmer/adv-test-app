package main

type Book struct {
  Title  string `json:"title"`
  Author  string `json:"author"`
  ISBN string `json:"isbn"`
  Description string `json:"description,omitempty"`
}

// This variable books contains a mapping of string type to the struct type Book with filled in values
// A few are filed in for this example.
var books = map[string]Book{
  "1234567890": Book{Title: "Flying on a plane", Author: "United Airlines", ISBN: "1234567890"},
  "0987654321": Book{Title: "The Cerner Interview ", Author: "Paul Cranmer ", ISBN: "0987654321"},
  "1111111111": Book{Title: "Snakes on a plane", Author: "Sam L. Jackson", ISBN: "1111111111"},
  "2222222222": Book{Title: "Typing on a plane", Author: "Bo Burnham", ISBN: "1234567890"},
}


// This function AllBooks() is a slice of books that returns all of the books stored in the variable books
func AllBooks() []Book {
  values := make([]Book, len(books))
  idx := 0
  for _, book := range books {
    values[idx] = book
    idx++
  }
  return values
}

// this allows the visitor to get the book by the isbn number

func GetBook(isbn string) (Book, bool) {
  book, found := books[isbn]
  return book, found
}

// creates a new book if one does not exist all ready
func CreateBook(book Book) (string, bool){
  _, exists := books[book.ISBN]
  if exists {
    return "", false
  }
  books[book.ISBN] = book
  return book.ISBN, true
}

// updates an existing books isbn...
func UpdateBook(isbn string, book Book) bool{
  _, exists := books[isbn]
  if exists {
    books[isbn] = book
  }
return exists
}


// delete removes a boof frome the map by the ISBN key
func DeleteBook(isbn string){
    delete(books, isbn)

}
