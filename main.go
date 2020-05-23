package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	Con "github.com/leiendrulat/goserv/Context"
	ServeFiles "github.com/leiendrulat/goserv/Handlers"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("starting server......")
	mux := http.NewServeMux()
	mux.HandleFunc("/home", ServeFiles.getHome)
	mux.HandleFunc("/bunsky", ServeFiles.getbunsky)

	//used to get other files css/js
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//used for keeping context and requests for the server for logging
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8082", Con.AddContext(c, handler)))

}
