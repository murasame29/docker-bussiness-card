package router

import (
	"net/http"

	"github.com/murasame29/docker-bussiness-card/internal/components"
	"github.com/murasame29/docker-bussiness-card/internal/handler"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type router struct {
	mux *http.ServeMux
}

func NewRouter() http.Handler {
	r := &router{
		mux: http.NewServeMux(),
	}

	r.appHandler()

	return r.mux
}

func (r *router) appHandler() {
	index := handler.NewIndex()
	r.mux.Handle("/", index)
	app.Route("/", components.NewIndex())
}
