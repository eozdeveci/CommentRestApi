package main

import (
	"net/http"

	"github.com/eozdeveci/CommentRestApi/internal/comment"
	"github.com/eozdeveci/CommentRestApi/internal/database"
	transportHTTP "github.com/eozdeveci/CommentRestApi/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

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
		log.Error("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "Comment Rest Api",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up Rest Api")
		log.Fatal(err)
	}
}
