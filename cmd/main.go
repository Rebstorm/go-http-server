package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	NameKey = "name"
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

type NameResponse struct {
	Names    []string `json:"names"`
	LastName string   `json:"lastName,omitempty"`
}

func middlewareFunc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Hello from middleware, Method" + request.Method)
		next.ServeHTTP(writer, request)
	})
}

func setupRoutes(router *http.ServeMux) {

	router.Handle("GET /", middlewareFunc(func(writer http.ResponseWriter, request *http.Request) {

		_, err := fmt.Fprint(writer, "Hello World!")
		if err != nil {
			panic(err)
		}
	}))

	router.HandleFunc("POST /{market}", func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			err := request.Body.Close()
			if err != nil {
				slog.Info("YOU DUN FUCKED UP CLOSING THE BODY, #YOLO")
			}

		}()

		respondersName := request.URL.Query()[NameKey]

		body := &NameResponse{}

		if err := json.NewDecoder(request.Body).Decode(body); err != nil {
			panic("YOLO")
		}

		if len(respondersName) < 1 {
			_, err := fmt.Fprint(writer, "You forgot to supply a name in a query parameter.")
			if err != nil {
				panic(err)
			}

			return
		}

		res := &NameResponse{
			Names: respondersName,
		}

		err := json.NewEncoder(writer).Encode(res)
		if err != nil {
			panic("yolo")
		}

	})
}
