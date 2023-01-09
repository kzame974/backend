package handler

//
//import (
//	"encoding/json"
//	"net/http"
//)
//
//func registerHandler(w http.ResponseWriter, r *http.Request) {
//	// décoder les données JSON du corps de la requête
//	var u utilisateur
//	err := json.NewDecoder(r.Body).Decode(&u)
//	if err != nil {
//		http.Error(w, "Données JSON non valides", http.StatusBadRequest)
//		return
//	}
//
//	// vérifier que le nom d'utilisateur n'est pas déjà utilisé
//	if _, ok := utilisateursEnregistres[u.user]; ok {
//		http.Error(w, "user déjà utilisé", http.StatusBadRequest)
//		return
//	}
//
//	// enregistrer l'utilisateur
//	utilisateursEnregistres[u.user] = u
//
//	// renvoyer un message de succès
//	w.Write([]byte("Utilisateur enregistré avec succès"))
//}
