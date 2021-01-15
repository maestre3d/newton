package controller

import "net/http"

// BookHTTP aggregate.Book HTTP endpoints
type BookHTTP struct {
}

func (h BookHTTP) create(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello there from book create endpoint"))
}
