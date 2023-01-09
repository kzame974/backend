package handler

//
//import (
//	"encoding/json"
//	"net/http"
//)
//
//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	// décoder les données JSON du corps de la requête
//	var u utilisateur
//	err := json.NewDecoder(r.Body).Decode(&u)
//	if err != nil {
//		http.Error(w, "Données JSON non valides", http.StatusBadRequest)
//		return
//	}
//
//	// vérifier que l'utilisateur est enregistré et que le mot de passe est correct
//	storedUser, ok := utilisateursEnregistres[u.user]
//	if !ok {
//		http.Error(w, "user ou mot de passe incorrect", http.StatusUnauthorized)
//		return
//	}
//	if storedUser.password != u.password {
//		http.Error(w, "user ou mot de passe incorrect", http.StatusUnauthorized)
//		return
//	}
//
//	// renvoyer un message de succès
//	w.Write([]byte("Connexion réussie"))
//}
