package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// BookHTTP aggregate.Book HTTP endpoints
type BookHTTP struct {
}

func (h BookHTTP) Route(r *mux.Router) {
	r.Path("/books").Methods(http.MethodPost).HandlerFunc(h.create)
}

func (h BookHTTP) create(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello there from book create endpoint"))
}
