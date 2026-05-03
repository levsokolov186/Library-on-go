package main

import (
	UI "Library/UI"
)

func main() {

	//user1 := &Authentication.User{}
	//user1.AddUser("Иван Петров", "male", "hgggg@gmail.com", "secure123", 25)
	//err := Authentication.AddUserToFile(*user1)
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println("✓ Пользователь 1 добавлен")
	//}
	//
	//emptyUser := Authentication.User{}
	//err = Authentication.AddUserToFile(emptyUser)
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println("✓ Пользователь 1 добавлен")
	//}

	//bookShelf := LibraryPackage.NewShelf[LibraryPackage.Book](1)
	//comicsShelf := LibraryPackage.NewShelf[LibraryPackage.Comics](2)
	//newspaperShelf := LibraryPackage.NewShelf[LibraryPackage.Newspaper](3)
	//
	//fmt.Println("Создали 3 полки:")
	//fmt.Printf("Полка для книг: номер=%d, элементов=%d\n", bookShelf.Number, bookShelf.Count())
	//fmt.Printf("Полка для комиксов: номер=%d, элементов=%d\n", comicsShelf.Number, comicsShelf.Count())
	//fmt.Printf("Полка для газет: номер=%d, элементов=%d\n", newspaperShelf.Number, newspaperShelf.Count())
	//
	//book1 := LibraryPackage.Book{}
	//book2 := LibraryPackage.Book{}
	//comics1 := LibraryPackage.Comics{}
	//newspaper1 := LibraryPackage.Newspaper{}
	//
	//if err := bookShelf.Add(book1); err != nil {
	//	log.Println(err)
	//}
	//if err := bookShelf.Add(book2); err != nil {
	//	log.Println(err)
	//}
	//if err := comicsShelf.Add(comics1); err != nil {
	//	log.Println(err)
	//}
	//if err := newspaperShelf.Add(newspaper1); err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println()
	//fmt.Println("После добавления объектов:")
	//fmt.Printf("На полке для книг теперь %d элемента(ов)\n", bookShelf.Count())
	//fmt.Printf("На полке для комиксов теперь %d элемента(ов)\n", comicsShelf.Count())
	//fmt.Printf("На полке для газет теперь %d элемента(ов)\n", newspaperShelf.Count())
	//
	//fmt.Println()
	//fmt.Println("Проверка заполненности:")
	//fmt.Printf("Полка книг заполнена? %t\n", bookShelf.IsFull())
	//fmt.Printf("Полка комиксов заполнена? %t\n", comicsShelf.IsFull())
	//fmt.Printf("Полка газет заполнена? %t\n", newspaperShelf.IsFull())

	UI.Entrance()
}
