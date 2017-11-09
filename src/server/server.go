package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	layoutPath := filepath.Join("templates", "layout.html")
	targetPath := filepath.Join("templates", "errorHandling.html")
	
	tmpl, _ := template.ParseFiles(layoutPath, targetPath)

	tmpl.ExecuteTemplate(w, "layout", nil)
}

func main() {
	//Must create a fileserver to the stc directory to access css folder
	// in stc.
	fileServer := http.FileServer(http.Dir("stc"))
	http.Handle("/stc/", http.StripPrefix("/stc/", fileServer))
	
	//Serve templates for url with matching "LetsGo"
	http.HandleFunc("/error", serveTemplate)
	http.ListenAndServe(":9090", nil)
}
