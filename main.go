package main

import (
	"go-practice/handlers"
	"log"
	"net/http"
	"os"
)

func main()  {
	l := log.New(os.Stdout, "go-practice", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
