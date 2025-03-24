package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {

	router := http.NewServeMux()

	fmt.Print("Starting HTTP Server...")

	setupRoutes(router)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic(err)
	}

}

func setupRoutes(router *http.ServeMux) {
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World!")
		if err != nil {
			panic(err)
		}
	})

	router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		respondersName := request.URL.Query().Get("name")

		if respondersName == "" {
			_, err := fmt.Fprint(writer, "You forgot to supply a name in a query parameter.")
			if err != nil {
				panic(err)
			}

			return
		}

		_, err := fmt.Fprint(writer, "Hello there "+respondersName+"!")
		if err != nil {
			panic(err)
		}

	})
}
