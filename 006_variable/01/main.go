package main

import (
	"log"
	"os"
	"text/template"
)
// $variable in tpl.gohtml assigns a value to a variable which can be accessed anywhere with {{variable}} inside the template
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
