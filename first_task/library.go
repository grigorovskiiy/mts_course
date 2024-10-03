package main

type Book struct {
	name      string
	author    string
	pages     int
	libraryId int
}

func (b *Book) GetAuthor() string {
	return b.author
}

type Library struct {
	storage Storage
	dict    map[string]int
}

func CreateLibrary(storage Storage) *Library {
	return &Library{storage, make(map[string]int)}
}

func (library *Library) GetBook(name string) Book {
	return library.storage.GetBook(library.dict[name])
}

func (library *Library) AddBook(book Book, Generator func(Book) int) {
	library.dict[book.name] = Generator(book)
	book.libraryId = library.dict[book.name]
	library.storage.AddBook(book, library.dict[book.name])

}

func FirstGenerator(book Book) int {
	result := 0
	for el := range book.name {
		result += el
	}
	for el := range book.author {
		result += el
	}
	return result

}

func SecondGenerator(book Book) int {
	result := 0
	for el := range book.name {
		result += el
	}
	for el := range book.author {
		result += el
	}
	result += book.pages
	return result
}
