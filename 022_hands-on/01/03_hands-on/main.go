package main;

import (
	"net/http"
	"html/template"
)

var tpl *template.Template;

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"));
}

func a (w http.ResponseWriter, req *http.Request ) {
	tpl.ExecuteTemplate(w, "index.gohtml", 42);
}

func main() {
	http.HandleFunc("/", a);
	http.ListenAndServe(":8080", nil);
}