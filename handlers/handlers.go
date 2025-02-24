package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Hero struct {
	Name            string
	BackgroundImage string
}

var hero Hero

// Define hero background images
var heroBackgrounds = map[string]string{
	"Adam Warlock":     "/static/assets/adam-warlock-background.jpg",
	"Black Panther":    "/static/assets/wakanda.jpg",
	"Black Widow":      "/static/assets/black-widow-background.jpg",
	"Captain America":  "/static/assets/captain-america-background.jpg",
	"Cloak and Dagger": "/static/assets/cloak-and-dagger-background.jpg",
	"Doctor Strange":   "/static/assets/doctor-strange-background.jpg",
	"Groot":            "/static/assets/groot-background.jpg",
	"Winter Soldier":   "/static/assets/winter-soldier-background.jpg",
	// Add more heroes here...
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func HeroNameHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData struct {
		Name string `json:"name"`
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	// Check if the name exists
	if requestData.Name == "" {
		http.Error(w, "Hero name is empty", http.StatusBadRequest)
		return
	}

	// Update the global hero variable
	hero.Name = requestData.Name
	hero.BackgroundImage = heroBackgrounds[hero.Name] // Fetch background from the map

	// Send response
	response := map[string]string{"message": "Hero updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HeroHandler(w http.ResponseWriter, r *http.Request) {
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
