package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	films :=[]Film{
		{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "The Dark Knight", Director: "Christopher Nolan"},
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		filmsModel := map[string][]Film{
			"Films": films,
		}

		tmpl.Execute(w, filmsModel)
	}

	addFilmHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		log.Println("Title: ", title)
		log.Println("Director: ", director)

		tmpl := template.Must(template.ParseFiles("index.html"))

		// Simulate a delay
		time.Sleep(1 * time.Second)

		film := Film{Title: title, Director: director}
		films = append(films, film)

		tmpl.ExecuteTemplate(w, "film-list-element", film)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add-film", addFilmHandler)

	// Set the port to align with the exposed port in the dev container
	port := 8000

	fmt.Println("Server is running on port: ", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
