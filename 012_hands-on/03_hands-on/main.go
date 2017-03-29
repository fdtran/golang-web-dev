
package main;

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template;

type hotel struct {
	Name, Address, City, Zip, Region string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"));
}

func main() {
 hotels := []hotel{
	 hotel{
		"Felix's Hotel",
		"123 Fake St.",
		"San Francisco",
		"94111",
		"Northern",
	 }, hotel {
		"Jonathan's Hotel",
		"123 Different St.",
		"San Mateo",
		"92221",
		"Northern",
	 },
 }

 err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", hotels);
 if err != nil {
	log.Fatalln(err);
 }

}