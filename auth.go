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

	fmt.Println("Try yo login")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user.Login, user.Password)

	err := db.QueryRow("SELECT id FROM users WHERE login = $1 AND password = $2", user.Login, user.Password).Scan(&user.ID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user.ID)

	if user.ID == 0 {
		fmt.Println(err.Error())
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

	fmt.Println(token)

	var signed_token Token
	signed_token.SignedToken, err = token.SignedString([]byte(signingKey))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(signed_token)

	// надо как-то протестировать печеньки
	//cookie := http.Cookie{Value: signed_token.SignedToken}
	cookie := http.Cookie{
		Name:     "Bearer",
		Value:    signed_token.SignedToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	fmt.Println(cookie)

	fmt.Println("Login")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Autorization", "Bearer "+signed_token.SignedToken)
	json.NewEncoder(w).Encode(signed_token)
}

// ???
func logout(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var l Link
	l.LinkExists = true

	fmt.Println("Link exists")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(l)
	w.WriteHeader(http.StatusOK)
}
