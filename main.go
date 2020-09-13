package main

import (
	"log"
	"net/http"
	"text/template"
)

type pizza struct {
	Ingredients string
	Crust       string
}

func main() {
	capreseTempl := template.Must(template.ParseFiles("template/caprese.html"))
	pepperoniTempl := template.Must(template.ParseFiles("template/pepperoni.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.Handle("/", http.HandlerFunc(root))

	http.Handle("/caprese", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		capresePizza := pizza{
			Ingredients: "tomato, basil, mozarella",
			Crust:       "thin",
		}
		capreseTempl.Execute(res, capresePizza)
	}))

	http.Handle("/pepperoni", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		pepperoniPizza := pizza{
			Ingredients: "pepperoni, cheese",
			Crust:       "thick",
		}
		pepperoniTempl.Execute(res, pepperoniPizza)
	}))

	log.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	rootTempl := []byte(
		`
		<html>
			<style>
				h1 { color: teal; }
			</style>
			<body>
				<h1>I am the root of all pizzas!"
			</body>
		</html>
		`,
	)

	res.Write(rootTempl)
}
