package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"todomodule/app"
	"todomodule/infrastructure/httphandlers"
	"todomodule/migrations"
	"todomodule/pkg/config"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/lib/pq"
	"todomodule/infrastructure/repos/dbrepo"
)

func main() {
	//connString2:="postgres://postgres:wizard@192.168.59.103:5432/postgres?sslmode=disable"
	fmt.Println("Start program")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err, "26")
	}
	//connString:="postgres://postgres:wizard@192.168.59.103:5432/postgres?sslmode=disable"
	runMigrate(cfg.DataSourceName)

	db, err := sqlx.Connect("postgres", cfg.DataSourceName)
	if err != nil {
		log.Fatalln(err, "33")
	}
	repo, err := dbrepo.NewDBRepo(db)
	if err != nil {
		log.Fatalln(err, "37")
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
		log.Fatal(err, "58")
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		log.Fatal(err, "62")
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No database migration applied")
		} else {
			log.Fatal(err, "68")
		}
	}
}
