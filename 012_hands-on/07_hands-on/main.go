package main;

import (
	"log"
	"text/template"
	"os"
)

var tpl *template.Template;

type item struct {
	Name string
	Cost int
}

type meal struct {
	Name string
	Options []item
}

type menu []meal;

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant;

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"));
}

func main() {
	r := restaurants{
		restaurant{
			"Felix's Restaurant",
			menu{
				meal{
					"Breakfast",
					[]item{
						item{
							"Omelette",
							10,
						},
						item {
							"Pancakes",
							8,
						}, 
						item {
							"Muffins",
							5,
						},
					},
				},
				meal {
					"Lunch",
					[]item{
						item{
							"Burger",
							12,
						}, 
						item {
							"Sandwich",
							9,
						},
						item {
							"Salmon",
							14,
						},
					},
				},
				meal {
					"Dinner",
					[]item{
						item{
							"Salad",
							18,
						}, 
						item {
							"Steak",
							15,
						}, 
						item {
							"Pasta",
							10,
						},
					},
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", r);
	if err != nil {
		log.Fatalln(err);
	}
}