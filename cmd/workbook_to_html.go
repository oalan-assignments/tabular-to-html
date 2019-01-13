package main

import (
	"fmt"
	"github.com/iCasComaasOzgunAlan/cmd/parser"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("../layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[1:]
		data, err := parser.ParseWorkbook("../" + fileName)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("<html><h1>%s</h1></html>", err)))
		} else {
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)

			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
