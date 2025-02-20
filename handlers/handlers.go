package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Hero struct {
	Name            string
	BackgroundImage string
}

func HeroHandler(w http.ResponseWriter, r *http.Request) {
	heroName := r.URL.Query().Get("name")
	if heroName == "" {
		http.Error(w, "Hero name is required", http.StatusBadRequest)
		return
	}

	// Define hero background images
	heroBackgrounds := map[string]string{
		"Adam Warlock":   "/assets/adam-right.jpg",
		"Winter Soldier": "/assets/winter.jpg",
		// Add more heroes here...
	}

	hero := Hero{
		Name:            heroName,
		BackgroundImage: heroBackgrounds[heroName],
	}

	// Render hero template
	tmpl, err := template.ParseFiles("templates/hero-template.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, hero)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
