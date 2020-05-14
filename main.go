package main

import (
	"database/sql"
	"fmt"
	"todomodule/domain"
	"todomodule/infrastructure/repos/memRepo"
)

// TODO Добавить UNIQUE в todo ---------
// TODO Добавить inmmemo repo с mutex и map --------
// TODO Убрать хардкод defaultConnStr, defaultDriverName. Добавить поле для коннекта в config.json
// TODO Создать package httphandlers в infractructure. В нем должен быть тип TodoHandlers struct. TodoHandlers
// будет обрабатывать http запросы. POST на CreateTodo и Get на GetAll. Нужно запустить сервер на localhost:8080

var defaultConnStr string = "user=postgres password=wizard dbname=postgres sslmode=disable"
var defaultDriverName string = "postgres"

func main() {
	fmt.Println("Start program")
	////ClearAllDb()
	//db, err := sqlx.Connect(defaultDriverName, defaultConnStr)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//repo, err := dbRepo.NewDbRepo(db)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//service := app.NewTodos(repo)
	//service.Create(domain.Todo{"Something first"})
	//service.Create(domain.Todo{"Anything else"})
	//
	//fmt.Println(service.GetAll())
	repo := memRepo.NewMemRepo()
	repo.Create(domain.Todo{"First"})
	repo.Create(domain.Todo{"Second"})
	fmt.Println(repo.GetAll())

}

func ClearAllDb() {
	connStr := "user=postgres password=wizard dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("delete from todos")
}
