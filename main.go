package main

import (
	"go-practice/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main()  {
	l := log.New(os.Stdout, "go-practice", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	s.ListenAndServe()
}
