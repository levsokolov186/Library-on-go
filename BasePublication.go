package main

type BasePublication struct {
	id              string // Уникальный идентификатор (например: "book_001", "comics_042")
	title           string // Название публикации
	publisher       string // Издатель
	publicationDate string // Дата публикации (формат: "2024-04-28")
	language        string // Язык публикации (например: "en", "ru", "fr")
	pageCount       int    // Общее количество страниц
	format          string // Формат (hardcover/softcover/digital/tabloid/broadsheet)
}
