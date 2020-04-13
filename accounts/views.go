package accounts

import (
	"encoding/json"
	"net/http"

	"../core"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello User")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	err := AccountsService.registerUser(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("User has been registered successfully.")
}

func login(w http.ResponseWriter, r *http.Request) {
	user, token, err := AccountsService.login(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	data := core.Serialize(UserLoginSerializer{}, user).(UserLoginSerializer)
	data.Token = token
	json.NewEncoder(w).Encode(data)
}
