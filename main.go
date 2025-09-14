package main

import (
	"embed"
	"log"

	"wails-todo-app/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/*
var assets embed.FS

func main() {

	todo := backend.NewTodoService()

	err := wails.Run(&options.App{
		Title:  "Wails ToDo",
		Width:  900,
		Height: 700,
		Assets: assets,
		Bind: []interface{}{
			todo,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
