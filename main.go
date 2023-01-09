package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// utilisateur est un type décrivant un utilisateur enregistré
type utilisateur struct {
	user     string `json:"nom d'utilisateur"`
	password string `json:"mot de passe"`
}

var utilisateursEnregistres = map[string]utilisateur{} // map stockant les utilisateurs enregistrés

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// décoder les données JSON du corps de la requête
	var u utilisateur
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Données JSON non valides", http.StatusBadRequest)
		return
	}

	// vérifier que le nom d'utilisateur n'est pas déjà utilisé
	if _, ok := utilisateursEnregistres[u.user]; ok {
		http.Error(w, "user déjà utilisé", http.StatusBadRequest)
		return
	}

	// enregistrer l'utilisateur
	utilisateursEnregistres[u.user] = u

	// renvoyer un message de succès
	w.Write([]byte("Utilisateur enregistré avec succès"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// décoder les données JSON du corps de la requête
	var u utilisateur
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Données JSON non valides", http.StatusBadRequest)
		return
	}

	// vérifier que l'utilisateur est enregistré et que le mot de passe est correct
	storedUser, ok := utilisateursEnregistres[u.user]
	if !ok {
		http.Error(w, "user ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}
	if storedUser.password != u.password {
		http.Error(w, "user ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	// renvoyer un message de succès
	w.Write([]byte("Connexion réussie"))
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	// démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
