package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow("SELECT id FROM users WHERE login = $1 AND password = $2", user.Login, user.Password).Scan(&user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		http.Error(w, err.Error(), 550)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			// токен действителен 2 часа, менять тут
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	var signed_token Token
	signed_token.SignedToken, err = token.SignedString([]byte(signingKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// надо как-то протестировать печеньки
	cookie := http.Cookie{Value: signed_token.SignedToken}
	http.SetCookie(w, &cookie)

	fmt.Println("Login")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(signed_token)
}

// ???
func logout(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var v string
	_ = db.QueryRow("SELECT id FROM users WHERE login = $1 AND password = $2", user.Login, user.Password).Scan(&v)
	if v != "" {
		http.Error(w, "StatusConflict", 409)
		return
	}

	_, err := db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Register")
	w.WriteHeader(http.StatusCreated)
}

func checkInvite(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)["link"]
	var id string
	err := db.QueryRow("SELECT id FROM invites WHERE link_text = $1", link).Scan(&id)
	if err != nil || id == "" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("Link exists")
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode()
}
