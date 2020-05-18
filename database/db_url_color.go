package database

import (
	st "../data_struct"
	"database/sql"
	_ "github.com/FogCreek/mini"
	_ "github.com/lib/pq"
	"log"
)

type UrlRepository struct {
	database *sql.DB
}

func ConnectDatabase(config *st.Config) (*sql.DB, error) {
	return sql.Open("postgres", config.DatabaseConnStr)
}

func (repository *UrlRepository) CreateTable() {
	repository.database.Exec("delete from img_Url_Color")
	_, err := repository.database.Exec("CREATE TABLE IF NOT EXISTS " +
		`img_Url_Color("id" integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,` +
		`"url_Name" varchar(100), "color" varchar(25))`)
	if err != nil {
		log.Print(err)
	} else {
		log.Println("Table created...")
	}
}

func (repository *UrlRepository) InsertUrl(urlColorBd []st.UrlImage) {
	for _, uC := range urlColorBd {
		repository.database.Exec("insert into img_Url_Color values(default, $1, $2)", uC.UrlImg, uC.Color)
	}
	log.Print("Insert OK")
}

func (repository *UrlRepository) FindAll() []st.UrlImage {
	rows, err := repository.database.Query("SELECT * FROM img_Url_Color")
	if err != nil {
		log.Print(err)
	} else {
		log.Print("Request passed...")
	}
	defer rows.Close()

	var urlImages []st.UrlImage

	for rows.Next() {
		var (
			id    int
			url   string
			color string
		)
		rows.Scan(&id, &url, &color)

		urlImages = append(urlImages, st.UrlImage{
			Id:     id,
			UrlImg: url,
			Color:  color,
		})
	}
	return urlImages
}

func NewUrlRepository(database *sql.DB) *UrlRepository {
	return &UrlRepository{database: database}
}
