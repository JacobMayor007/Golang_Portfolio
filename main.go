package main

import (
	"fmt"
	"log"
	"myportfolio/routes"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/projects", routes.Projects)
	http.HandleFunc("/about", routes.About)

	fmt.Println("🚀 Server is running on http://localhost:21123")
	err := http.ListenAndServe(":21123", nil)
	if err != nil {
		log.Fatal(err)
	}
}
