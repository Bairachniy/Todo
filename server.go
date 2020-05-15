package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"todomodule/app"
	"todomodule/domain"
	"todomodule/infrastructure/httphandlers"
	"todomodule/infrastructure/repos/dbRepo"
	"todomodule/pkg/config"
)

func main() {
	fmt.Println("Start program")
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := dbRepo.NewDbRepo(db)
	if err != nil {
		log.Fatalln(err)
	}
	service := app.NewTodos(repo)
	handle := httphandlers.NewHttpHandler(service)
	handle.Create(domain.Todo{"Jopa"})
}
