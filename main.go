package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rs/cors"

)

//used to get template stuff
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

var err error

func main() {

	fmt.Println("starting server......")
	mux := http.NewServeMux()
	mux.HandleFunc("/home", getHome)

	//used to get other files css/js
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//used for keeping context and requests for the server for logging
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Con.AddContext(c, handler)))

}

//template to get file
func getHome(w http.ResponseWriter, r *http.Request) {

	//switch statement for get or post
	switch r.Method {

	case "GET":
		//go get html file
		err := tpl.ExecuteTemplate(w, "mysite.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}

	case "POST":
		//go get html file
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
