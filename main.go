package main

import (
	st "./data_struct"
	db "./database"
	srv "./server"
	s "./service"
	"go.uber.org/dig"
	"log"
)

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *srv.Server) {
		server.RunUrlColor()
		server.Run()
	})
	if err != nil {
		log.Print(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(st.NewConfig)
	container.Provide(db.ConnectDatabase)
	container.Provide(s.NewUrlColorService)
	container.Provide(db.NewUrlRepository)
	container.Provide(srv.NewServer)

	return container
}
