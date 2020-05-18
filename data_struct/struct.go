package data_struct

type UrlImage struct {
	Id     int
	UrlImg string
	Color  string
}

type Config struct {
	Enabled         bool
	DatabaseConnStr string
	Port            string
	UrlString       string
}

func NewConfig() *Config {
	return &Config{
		Enabled:         true, //будет ли наше приложение возвращать реальные данные.
		DatabaseConnStr: "user=postgres password=123 dbname=mydb sslmode=disable",
		Port:            "8181",
	}
}
