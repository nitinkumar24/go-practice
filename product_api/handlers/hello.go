package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("hello world")

	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "Hello %s\n", d)
}