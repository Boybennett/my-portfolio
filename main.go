package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Static Files (Images, CSS, etc.)
	// URL: localhost:8000/static/logo.png -> Folder: ./static/logo.png
	staticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))

	// 2. The Home Page (Root)
	// URL: localhost:8000/ -> File: ./public/dashboard.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We check if path is exactly "/" so that /random-page doesn't show the dashboard
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./public/dashboard.html")
			return
		}
		// If they type a URL we haven't defined, show a 404
		http.NotFound(w, r)
	})

	// 3. The About Page (Clean URL)
	// URL: localhost:8000/about -> File: ./public/about.html

	// 4. Other Pages
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/dashboard.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/about.html")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/contact.html")
	})

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/projects.html")
	})

	http.HandleFunc("/tech-stack", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/tech-stack.html")
	})

	fmt.Println("Server starting on http://localhost:8000")
	fmt.Println("Dashboard: http://localhost:8000/")
	fmt.Println("About:     http://localhost:8000/about")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}