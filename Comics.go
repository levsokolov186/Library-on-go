package main

type Comics struct {
	BasePublication
	issueNumber     int      // Номер выпуска (например: 1, 42, 100)
	volume          int      // Том (volume) комикса
	series          string   // Название серии (например: "Batman", "One Piece", "Saga")
	writers         []string // Сценаристы (кто написал сюжет и диалоги)
	artists         []string // Художники (кто рисовал основные карандашные рисунки)
	inkers          []string // Контуровщики (обводка карандашного рисунка тушью)
	colorists       []string // Колористы (заливка цветом)
	letterers       []string // Леттереры (размещение текста, шрифтов, облаков диалога)
	coverArtists    []string // Художники обложки
	genre           string   // Жанр (например: "Superhero", "Manga", "Horror", "Noir")
	ageRating       string   // Возрастной рейтинг ("All Ages", "Teen", "Mature", "Explicit")
	synopsis        string   // Краткое описание сюжета (синопсис)
	hasVariantCover bool     // Наличие вариантов обложки (variant covers)
}
