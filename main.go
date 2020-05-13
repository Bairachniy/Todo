package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"todomodule/app"
	"todomodule/domain"
	"todomodule/infrastructure/repos/dbRepo"
)

// TODO Добавить UNIQUE в todo
// TODO Добавить inmmemo repo с mutex и map
// TODO Убрать хардкод defaultConnStr, defaultDriverName. Добавить поле для коннекта в config.json
// TODO Создать package httphandlers в infractructure. В нем должен быть тип TodoHandlers struct. TodoHandlers
// будет обрабатывать http запросы. POST на CreateTodo и Get на GetAll. Нужно запустить сервер на localhost:8080

var defaultConnStr string = "user=postgres password=Qaz54321 dbname=testdb sslmode=disable"
var defaultDriverName string = "postgres"

func main() {
	fmt.Println("Start program")
	db, err := sqlx.Connect(defaultDriverName, defaultConnStr)
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := dbRepo.NewDbRepo(db)
	if err != nil {
		log.Fatalln(err)
	}
	service := app.NewTodos(repo)
	service.Create(domain.Todo{"Something first"})
	service.Create(domain.Todo{"Anything else"})
	fmt.Println(service.GetAll())
}
