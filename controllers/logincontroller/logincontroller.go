package logincontroller

import (
	"encoding/json"
	"go-web-native/entities"
	"go-web-native/middleware"
	"io"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	w.Header().Set("Content-Type", "application/json")
	var user entities.Users
	defer r.Body.Close()
	request, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.Unmarshal(request, &user)

	// fmt.Println(r.Method)
	// if r.Method != "POST" {
	// 	http.Error(w, "Unsupported http method", http.StatusBadRequest)
	// 	return
	// }

	// username, password, ok := r.BasicAuth()
	// if !ok {
	// 	http.Error(w, "Invalid username or password", http.StatusBadRequest)
	// 	return
	// }

	token, err := middleware.CreateToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"token": token}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}
