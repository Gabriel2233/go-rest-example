package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mural-app/server/infra/repository"
	"github.com/mural-app/server/model"
)

type WebServer struct {
	Repo repository.ThoughtRepositoryDb
}

func NewServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()

	e.GET("/thoughts", w.getAll)
	e.GET("/thoughts-by-tag", w.getByTag)
	e.POST("/create", w.createThought)

	e.Logger.Fatal(e.Start(":8080"))
}

func (w WebServer) getAll(e echo.Context) error {
	data := w.Repo.FindAll()

	return e.JSON(http.StatusOK, data)
}

func (w WebServer) getByTag(e echo.Context) error {
	tag := e.QueryParam("tag")

	data := w.Repo.FindByTag(tag)

	return e.JSON(http.StatusOK, data)
}

func (w WebServer) createThought(e echo.Context) error {
	newThought := model.NewThought()

	if err := e.Bind(newThought); err != nil {
		return err
	}

	w.Repo.Add(newThought)

	return e.JSON(http.StatusCreated, newThought)
}
