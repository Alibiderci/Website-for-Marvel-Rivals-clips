package main

import (
	"fmt"
	"net/http"

	"github.com/Alibiderci/website-for-clips/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("hero/", handlers.HeroHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
