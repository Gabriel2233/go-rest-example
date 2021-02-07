package main

import (
	"github.com/mural-app/server/http/rest"
	"github.com/mural-app/server/infra/db"
	"github.com/mural-app/server/infra/repository"
)

func main() {

	database := db.ConnectDB()

	repo := repository.ThoughtRepositoryDb{
		Db: database,
	}

	webserver := rest.NewServer()
	webserver.Repo = repo
	webserver.Serve()
}
