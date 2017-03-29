package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	//returns a pointer to a template i.e. container with all our parsed Data.
	//ParseFiles can take any number of strings
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	//execute takes writer and data from docs
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
