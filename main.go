package main

import (
	"fmt"
	"net/http"
)

func main() {
	// démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
