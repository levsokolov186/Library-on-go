package main

type Newspaper struct {
	BasePublication
	name        string   // Полное название газеты (например: "The New York Times", "The Guardian")
	editionType string   // Тип выпуска ("daily" - ежедневная, "sunday" - воскресная, "weekly" - еженедельная, "weekend" - уикенд)
	volume      int      // Год/том издания (например: 174 - 174-й год издания)
	issueNumber int      // Номер выпуска (сквозной или внутри года)
	theme       string   // Основная тематика номера (например: "Politics", "Economy", "Sports")
	coverStory  string   // Главная тема/новость номера (например: "Presidential Election", "Climate Summit")
	issn        string   // Международный стандартный номер газеты (ISSN)
	circulation int      // Тираж (количество напечатанных экземпляров)
	sections    []string // Список секций/разделов (например: ["News", "Opinion", "Business", "Arts"])
	price       string   // Цена экземпляра (например: "$3.00", "€2.50", "£2.00")
}
