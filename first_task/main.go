package main

import "fmt"

func main() {

	var mas = []Book{{"Война и мир", "Л.Н. Толстой", 1300, 0}, {"Мертвые души", "Н.В. Гоголь", 352, 0}}
	library := CreateLibrary(CreateSlice(make([]Book, 0)))

	for _, book := range mas {
		library.AddBook(book, FirstGenerator)
	}

	fmt.Println(library.GetBook("Война и мир"))
	fmt.Println(library.GetBook("Мертвые души"))

	library.storage = CreateMap(make(map[int]Book))
	mas = []Book{{"Преступление и наказание", "Ф.М. Достоевский", 496, 0}}

	for _, book := range mas {
		library.AddBook(book, SecondGenerator)
	}

	fmt.Println(library.GetBook("Преступление и наказание"))
}
