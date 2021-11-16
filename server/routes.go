package server

import (
	"bytes"
	"github.com/Oppodelldog/go-simple-webapp-template/assets"
	"github.com/Oppodelldog/go-simple-webapp-template/page"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

const RouteIndex = "/index"

func NewRouter() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(RouteIndex, indexHandler).Methods(http.MethodGet)

	m.PathPrefix("/css/").Handler(http.FileServer(http.FS(assets.Styles))).Methods(http.MethodGet)
	m.PathPrefix("/icons/").Handler(http.FileServer(http.FS(assets.Icons))).Methods(http.MethodGet)
	m.PathPrefix("/js/").Handler(http.FileServer(http.FS(assets.Javascript))).Methods(http.MethodGet)

	m.PathPrefix("/").HandlerFunc(catchAll).Methods(http.MethodGet)
	return m
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	writeBuffer := bytes.NewBufferString("")
	err := page.RenderIndexPage(writeBuffer, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error while rendering index page: %v", err)
		return
	}

	_, err = writeBuffer.WriteTo(w)
	if err != nil {
		logrus.Errorf("error while writing page content to client: %v", err)
	}
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, RouteIndex, http.StatusFound)
}
