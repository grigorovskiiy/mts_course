package main

type Storage interface {
	GetBook(id int) Book
	AddBook(book Book, id int)
}

type Map struct {
	storage map[int]Book
}

func CreateMap(storage map[int]Book) *Map {
	return &Map{storage: storage}
}

func (m *Map) GetBook(id int) Book {
	return m.storage[id]
}

func (m *Map) AddBook(book Book, id int) {
	m.storage[id] = book
}

type Slice struct {
	storage []Book
}

func CreateSlice(storage []Book) *Slice {
	return &Slice{storage}
}

func (s *Slice) GetBook(id int) Book {
	i := 0
	for ind, book := range s.storage {
		if book.libraryId == id {
			i = ind
		}
	}
	return s.storage[i]
}

func (s *Slice) AddBook(book Book, id int) {
	s.storage = append(s.storage, book)
}
