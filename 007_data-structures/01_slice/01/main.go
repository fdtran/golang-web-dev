package main

import (
	"log"
	"os"
	"text/template"
)
//"go fmt ./..." formats all of the go code in directory
//range in tpl.gohtml iterates over a slice and requires an {{end}}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
