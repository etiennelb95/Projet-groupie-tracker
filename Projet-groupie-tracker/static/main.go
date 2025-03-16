package main

import (
	"html/template"
	"net/http"
)

// Structure pour représenter un anime
type Anime struct {
	Title string
	Score float64
	Image string
}

// Fonction pour récupérer les animes depuis ton API (exemple statique ici)
func getTopAnime() []Anime {
	return []Anime{
		{"Attack on Titan", 9.1, "https://cdn.myanimelist.net/images/anime/10/47347.jpg"},
		{"Death Note", 8.7, "https://cdn.myanimelist.net/images/anime/9/9453.jpg"},
		{"Naruto", 7.9, "https://cdn.myanimelist.net/images/anime/13/17405.jpg"},
	}
}

// Handler pour la page d'accueil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	animes := getTopAnime() // Remplace par un appel API réel si nécessaire
	tmpl.Execute(w, animes)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}
