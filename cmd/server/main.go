package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up APP")
	return nil
}

func main() {
	fmt.Println("Go - Comment Rest Api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up Rest Api")
		fmt.Println(err)
	}
}
