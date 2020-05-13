package services

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"module/infrastructure/repos"
//)
//
////type Todos struct {
////	repo repos.Repo
////}
//type TodoTmp struct {
//	id   int
//	todo string
//}
//
//
////func NewTodo(repo Repo) *Todo{
////	return &Todo{repo: repo}
////}
////func NewTodos(repo repos.Repo) *Todos {
////	return &Todos{repo: repo}
////}
////func (t *Todos) Create(todo string) error {
////	return t.repo.Create(todo)
////}
////func (t *Todos) GetAll() ([]string,error) {
////	return t.repo.GetAll()
////}
//
//type ToBase struct {
//	db sql.DB
//}
//type ToLocal struct {
//	todo []string
//}
//
//func (todo *ToLocal) Create(s string) error {
//	todo.todo = append(todo.todo, s)
//	return nil
//}
////func (todo *ToBase) Create(s string) error {
////	ifExist := true
////	for _, t:= range GetAll() {
////		if s == t {
////			ifExist = false
////			fmt.Println("Such todo is already exist in todoList")
////			break
////		}
////	}
////	if ifExist {
////		connStr := "user=postgres password=wizard dbname=postgres sslmode=disable"
////		db, err := sql.Open("postgres", connStr)
////		if err != nil {
////			panic(err)
////		}
////		defer db.Close()
////
////		db.Exec("insert into todos (todo) values ($1)", s)
////		if err != nil {
////			panic(err)
////		}
////		fmt.Println("Created: ", s)
////	}
////	return nil
////}
//
//func (todo *ToLocal) GetAll() ([]string,error) {
//	return todo.todo,nil
//}
//func (todo *ToBase) GetAll() ([]string,error) {
//	connStr := "user=postgres password=wizard dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	rows, err := db.Query("select * from todos")
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		t := TodoTmp{}
//		err := rows.Scan(&t.id, &t.todo)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		todo.todo = append(todo.todo, t.todo)
//	}
//	return todo.todo,nil
//}
//
//func (todo *Todos) ClearAllDb() {
//	connStr := "user=postgres password=wizard dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//	db.Exec("delete from todos")
//}
//
//func GetAll() ([]string) {
//	connStr := "user=postgres password=wizard dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	rows, err := db.Query("select * from todos")
//	if err != nil {
//		panic(err)
//	}
//	var tmp []string
//	for rows.Next() {
//		t := TodoTmp{}
//		err := rows.Scan(&t.id, &t.todo)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		tmp = append(tmp, t.todo)
//	}
//	return tmp
//}
