package main

import "fmt"

type Iterator interface {
    HasNext() bool
    Next() *Book
}


type Book struct {
    Title  string
    Author string
}

type BookCollection struct {
    books []*Book
}

func (c *BookCollection) AddBook(book *Book) {
    c.books = append(c.books, book)
}

func (c *BookCollection) CreateIterator() Iterator {
    return &BookIterator{
        collection: c,
        index:      0,
    }
}


type BookIterator struct {
    collection *BookCollection
    index      int
}

func (it *BookIterator) HasNext() bool {
    return it.index < len(it.collection.books)
}

func (it *BookIterator) Next() *Book {
    if !it.HasNext() {
        return nil
    }
    book := it.collection.books[it.index]
    it.index++
    return book
}


func main() {
    collection := &BookCollection{}
    collection.AddBook(&Book{Title: "Go Programming", Author: "Robert Griesemer"})
    collection.AddBook(&Book{Title: "Patterns in Go", Author: "Maxi"})
    collection.AddBook(&Book{Title: "Advanced Go", Author: "Mini"})

    iterator := collection.CreateIterator()
    for iterator.HasNext() {
        book := iterator.Next()
        fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
    }
}
