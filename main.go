package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todomodule/app"
	"todomodule/infrastructure/repos/dbRepo"
)

func main() {
	fmt.Println("Start program")
	db := &sqlx.DB{}
	repo := dbRepo.NewDbRepo(db)
	service := app.NewTodos(repo)
	//service.Create(domain.Todo{"Something first"})
	//service.Create(domain.Todo{"Anything else"})
	fmt.Println(service.GetAll())
}
