package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"todomodule/app"
	"todomodule/infrastructure/httphandlers"

	//"github.com/labstack/echo/middleware"
	"log"

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
	//uptodo:=app.UpdateTodoCommand{ID: "81",Name: "udated"}
	//if err != nil {
	//	log.Fatalln(err)
	//}
	dbrepo := app.NewTodos(repo)
	dbrepo.Delete("81")
	fmt.Println(dbrepo.GetAll())

	service := app.NewTodos(repo)
	handle := httphandlers.NewHttpHandler(service)
	e := echo.New()
	e.POST("/todos", handle.CreateTodo)
	e.POST("/todos/:id", handle.UpdateTodo)
	e.GET("/todos", handle.GetAllTodo)
	e.DELETE("/todos/:name", handle.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))

}
