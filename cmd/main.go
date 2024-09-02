package main

import "github.com/eagle7802/todo-list/internal/app"

func main() {
	if err := app.Start(); err != nil {
		return
	}

}
