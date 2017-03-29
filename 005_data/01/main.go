package main

import (
	"log"
	"os"
	"text/template"
)
// "." is reference to the current piece of data in tp.gohtml i.e. 42. third argument is the data
//only get to pass one piece of data, but can pass in struct (objects);
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
