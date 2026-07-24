package ShelvesUI

import (
	LibraryPackage "Library/Library"
	"fmt"
	"log"
)

func PrintShelves() {
	bookShelfOne := LibraryPackage.NewShelf[LibraryPackage.Book](1)
	bookShelfTwo := LibraryPackage.NewShelf[LibraryPackage.Book](2)
	comicsShelfOne := LibraryPackage.NewShelf[LibraryPackage.Comics](3)
	comicsShelfTwo := LibraryPackage.NewShelf[LibraryPackage.Comics](4)
	newspaperShelf := LibraryPackage.NewShelf[LibraryPackage.Newspaper](5)

	fillBooks(bookShelfOne, []LibraryPackage.Book{
		LibraryPackage.NewBook("book-001", "Мастер и Маргарита", "Михаил Булгаков", 480),
		LibraryPackage.NewBook("book-002", "Преступление и наказание", "Фёдор Достоевский", 672),
		LibraryPackage.NewBook("book-003", "Война и мир", "Лев Толстой", 1225),
		LibraryPackage.NewBook("book-004", "Отцы и дети", "Иван Тургенев", 288),
		LibraryPackage.NewBook("book-005", "Мы", "Евгений Замятин", 256),
	})

	fillBooks(bookShelfTwo, []LibraryPackage.Book{
		LibraryPackage.NewBook("book-006", "1984", "Джордж Оруэлл", 328),
		LibraryPackage.NewBook("book-007", "Дюна", "Фрэнк Герберт", 544),
		LibraryPackage.NewBook("book-008", "Солярис", "Станислав Лем", 288),
		LibraryPackage.NewBook("book-009", "Над пропастью во ржи", "Джером Сэлинджер", 224),
		LibraryPackage.NewBook("book-010", "Три товарища", "Эрих Мария Ремарк", 480),
	})

	fillComics(comicsShelfOne, []LibraryPackage.Comics{
		LibraryPackage.NewComics("comics-001", "Batman: Year One", "Batman", 1),
		LibraryPackage.NewComics("comics-002", "Spider-Man: Blue", "Spider-Man", 2),
		LibraryPackage.NewComics("comics-003", "Saga", "Saga", 3),
		LibraryPackage.NewComics("comics-004", "Hellboy", "Hellboy", 4),
		LibraryPackage.NewComics("comics-005", "Watchmen", "Watchmen", 5),
	})

	fillComics(comicsShelfTwo, []LibraryPackage.Comics{
		LibraryPackage.NewComics("comics-006", "One Piece", "One Piece", 101),
		LibraryPackage.NewComics("comics-007", "Naruto", "Naruto", 58),
		LibraryPackage.NewComics("comics-008", "Bleach", "Bleach", 33),
		LibraryPackage.NewComics("comics-009", "Berserk", "Berserk", 17),
		LibraryPackage.NewComics("comics-010", "Fullmetal Alchemist", "Fullmetal Alchemist", 12),
	})

	fillNewspapers(newspaperShelf, []LibraryPackage.Newspaper{
		LibraryPackage.NewNewspaper("news-001", "Вестник библиотеки", "Культура", 1),
		LibraryPackage.NewNewspaper("news-002", "Городские новости", "Общество", 2),
		LibraryPackage.NewNewspaper("news-003", "Научный обзор", "Наука", 3),
		LibraryPackage.NewNewspaper("news-004", "Спортивный курьер", "Спорт", 4),
		LibraryPackage.NewNewspaper("news-005", "Экономический день", "Экономика", 5),
	})

	fmt.Println("Создали 5 полностью заполненных полок:")
	fmt.Println()

	printBookShelf("Полка 1: книги", bookShelfOne)
	printBookShelf("Полка 2: книги", bookShelfTwo)
	printComicsShelf("Полка 3: комиксы", comicsShelfOne)
	printComicsShelf("Полка 4: комиксы", comicsShelfTwo)
	printNewspaperShelf("Полка 5: газеты", newspaperShelf)
}

func fillBooks(shelf *LibraryPackage.Shelf[LibraryPackage.Book], items []LibraryPackage.Book) {
	for _, item := range items {
		if err := shelf.Add(item); err != nil {
			log.Println(err)
		}
	}
}

func fillComics(shelf *LibraryPackage.Shelf[LibraryPackage.Comics], items []LibraryPackage.Comics) {
	for _, item := range items {
		if err := shelf.Add(item); err != nil {
			log.Println(err)
		}
	}
}

func fillNewspapers(shelf *LibraryPackage.Shelf[LibraryPackage.Newspaper], items []LibraryPackage.Newspaper) {
	for _, item := range items {
		if err := shelf.Add(item); err != nil {
			log.Println(err)
		}
	}
}

func printBookShelf(label string, shelf *LibraryPackage.Shelf[LibraryPackage.Book]) {
	fmt.Printf("%s\n", label)
	fmt.Printf("Номер полки: %d\n", shelf.Number)
	fmt.Printf("Элементов: %d из %d\n", shelf.Count(), LibraryPackage.MaxItemsPerShelf)
	fmt.Printf("Заполнена: %t\n", shelf.IsFull())
	for index, item := range shelf.Items {
		fmt.Printf("%d. %s\n", index+1, item.Summary())
	}
	fmt.Println()
}

func printComicsShelf(label string, shelf *LibraryPackage.Shelf[LibraryPackage.Comics]) {
	fmt.Printf("%s\n", label)
	fmt.Printf("Номер полки: %d\n", shelf.Number)
	fmt.Printf("Элементов: %d из %d\n", shelf.Count(), LibraryPackage.MaxItemsPerShelf)
	fmt.Printf("Заполнена: %t\n", shelf.IsFull())
	for index, item := range shelf.Items {
		fmt.Printf("%d. %s\n", index+1, item.Summary())
	}
	fmt.Println()
}

func printNewspaperShelf(label string, shelf *LibraryPackage.Shelf[LibraryPackage.Newspaper]) {
	fmt.Printf("%s\n", label)
	fmt.Printf("Номер полки: %d\n", shelf.Number)
	fmt.Printf("Элементов: %d из %d\n", shelf.Count(), LibraryPackage.MaxItemsPerShelf)
	fmt.Printf("Заполнена: %t\n", shelf.IsFull())
	for index, item := range shelf.Items {
		fmt.Printf("%d. %s\n", index+1, item.Summary())
	}
	fmt.Println()
}
