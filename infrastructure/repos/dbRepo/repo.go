package dbRepo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
	"todomodule/app"
	"todomodule/domain"
)

var schema = `CREATE TABLE if not exists todos (
id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
todo  text NOT NULL,
UNIQUE(todo)
)`

var makeunique = `ALTER TABLE todos
ADD CONSTRAINT todo UNIQUE (todo);`

type dbRepo struct {
	db *sqlx.DB
}

func (r dbRepo) Update(todo domain.Todo) error {
	_, err := r.db.Exec(`
	UPDATE todos
	SET todo=$1
	WHERE
	id=$2;`, todo.Name(), todo.ID())
	return err
}

func (r dbRepo) Delete(id string) error {
	_, err := r.db.Exec(`
	DELETE FROM todos WHERE id=($1)`, id)
	return err
}

func (r dbRepo) MakeUniqueTodoQuery() error {
	_, err := r.db.Exec(makeunique)
	return err
}
func (r dbRepo) Create(todo domain.Todo) error {
	_, err := r.db.Exec(`
	INSERT INTO todos (todo) VALUES ($1) on conflict do nothing`, todo.Name())
	return err
}
func (r dbRepo) GetAll() ([]domain.Todo, error) {
	var todos []Todo
	err := r.db.Select(&todos, "select * from todos")
	domainTodos := make([]domain.Todo, 0, len(todos))
	for i := range todos {
		tmp, err := domain.ParseTodo(strconv.Itoa(todos[i].ID), todos[i].Name)
		if err != nil {
			panic(err)
		}
		domainTodos = append(domainTodos, tmp)
	}
	return domainTodos, err
}

func NewDbRepo(db *sqlx.DB) (app.Repo, error) {
	_, err := db.Exec(schema)
	return &dbRepo{db: db}, err
}
