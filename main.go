package main

import (
	"log"
	"net/http"
)

func main() {

	// Fichiers statiques (images, css, etc.)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/",
		fs,
	)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
