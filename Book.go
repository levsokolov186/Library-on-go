package main

type Book struct {
	BasePublication
	size        int      // Размер/формат книги в мм (например: 150x220)
	author      string   // Автор или соавторы (основной автор)
	authors     []string // Список всех авторов (если их несколько)
	isbn        string   // Международный стандартный книжный номер (ISBN-10 или ISBN-13)
	genre       string   // Жанр книги (например: "Fantasy", "Science Fiction")
	edition     int      // Номер издания (1 = первое издание)
	series      string   // Название серии, если книга входит в серию (например: "The Witcher")
	seriesIndex int      // Порядковый номер книги в серии (например: 3 - третья книга)
}
