package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(writer, "Hello %s\n", d)
	})

	http.HandleFunc("/goodBye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("good bye")
	})

	http.ListenAndServe(":9090", nil)
}
