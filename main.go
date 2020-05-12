package main

import (

	"fmt"
	"module/services"
)

func main(){
	fmt.Println("Start program")
	var toBase services.Repo=new(services.ToBase)
	var toLocal services.Repo=new(services.ToLocal)
	todoBase:=services.NewTodos(toBase)
	todoLocal:=services.NewTodos(toLocal)
	todoBase.Create("First to Base")
	todoBase.Create("Second to Base")
	todoLocal.Create("First to Local")
	todoLocal.Create("Second to Local")
	fmt.Println(todoBase.GetAll())
	fmt.Println(todoLocal.GetAll())
}