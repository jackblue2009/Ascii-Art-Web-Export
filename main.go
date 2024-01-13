package main

import (
	server "ascii-art-web/pkg/server"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.PathHandler)
	//Below is address for local host server
	link := "http://localhost:8080"
	fmt.Println("\033[36mServer Connected...\033[0m")
	fmt.Printf("\033[36mlink on: %s\033[0m\n", link)
	http.Handle("/css/", http.FileServer(http.Dir("templates")))
	http.Handle("/img/", http.FileServer(http.Dir("templates")))
	http.Handle("/js/", http.FileServer(http.Dir("templates")))
	http.ListenAndServe(":8080", nil)
}
