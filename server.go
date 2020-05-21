package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"todomodule/app"
	"todomodule/infrastructure/httphandlers"

	"log"

	"todomodule/infrastructure/repos/dbrepo"
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
	repo, err := dbrepo.NewDbRepo(db)
	if err != nil {
		log.Fatalln(err)
	}

	dbRepo := app.NewTodos(repo)

	fmt.Println(dbRepo.GetAll())

	service := app.NewTodos(repo)
	handle := httphandlers.NewHTTPHandler(service)
	e := echo.New()
	e.POST("/todos", handle.CreateTodo)
	e.POST("/todos/:id", handle.UpdateTodo)
	e.GET("/todos", handle.GetAllTodo)
	e.DELETE("/todos/:id", handle.DeleteTodo)
	e.Logger.Fatal(e.Start(":8070"))
}
