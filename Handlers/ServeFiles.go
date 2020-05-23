package ServeFiles

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	Headers "github.com/leiendrulat/goserv/Handlers/Headers"
)

//used to get template stuff
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

var err error

//template to get file
func getHome(w http.ResponseWriter, r *http.Request) {
	Headers.Header(w, r)
	//switch statement for get or post
	switch r.Method {

	case "GET":
		//go get html file
		err := tpl.ExecuteTemplate(w, "mysite.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
func getbunsky(w http.ResponseWriter, r *http.Request) {
	Headers.Header(w, r)

	//switch statement for get or post
	switch r.Method {

	case "GET":
		//go get html file
		err := tpl.ExecuteTemplate(w, "bunsky.html", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
