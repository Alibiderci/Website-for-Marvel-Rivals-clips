package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alibiderci/website-for-clips/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomePageHandler)

	http.HandleFunc("POST /heroName", handlers.HeroNameHandler)
	http.HandleFunc("GET /hero", handlers.HeroHandler)

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
