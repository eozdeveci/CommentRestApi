package main

import (
	"fmt"
	"net/http"

	"github.com/eozdeveci/CommentRestApi/internal/comment"
	"github.com/eozdeveci/CommentRestApi/internal/database"
	transportHTTP "github.com/eozdeveci/CommentRestApi/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up APP")

	var err error

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
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
