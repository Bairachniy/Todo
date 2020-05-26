package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"todomodule/app"
	"todomodule/infrastructure/httphandlers"
	"todomodule/migrations"

	"log"

	_ "github.com/lib/pq"
	"todomodule/infrastructure/repos/dbrepo"
	"todomodule/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
)

func main() {

	fmt.Println("Start program")
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	runMigrate(cfg.DataSourceName)

	db, err := sqlx.Connect(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := dbrepo.NewDBRepo(db)
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
	e.Logger.Fatal(e.Start(":8080"))
}

func runMigrate(dsn string) {
	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No database migration applied")
		} else {
			log.Fatal(err)
		}
	}
}
