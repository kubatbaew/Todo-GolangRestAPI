package main

import (
	"github.com/zhashkevych/todo-app"
	handler2 "github.com/zhashkevych/todo-app/package/handler"
	"log"
)

func main() {
	handlers := new(handler2.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
