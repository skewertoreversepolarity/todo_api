package main

import (
	"log"
	todo "todo-app"
	handler2 "todo-app/pkg/handler"
)

func main() {
	handler := new(handler2.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occoring while runing http server %s", err)
	}
}
