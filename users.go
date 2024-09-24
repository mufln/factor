package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	rows, err := db.Query("SELECT u.id,u.login,u.rights_level,e.name,e.profile_pic_path FROM users AS u INNER JOIN employees AS e ON u.id = e.id;")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Login, &user.RightsLevel, &user.Name, &user.ProfilePicPath); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	fmt.Println("Get users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := db.Exec("INSERT INTO users (login) VALUES ($1)", user.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Create user")
	w.WriteHeader(http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Login != "" {
		_, err := db.Exec("UPDATE users SET login = $1 WHERE id = $2", user.Login, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if user.Name != "" {
		_, err := db.Exec("UPDATE employees SET name = $1 WHERE id = $2", user.Name, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if user.Password != "" {
		_, err := db.Exec("UPDATE users SET password = $1 WHERE id = $2", user.Password, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
	if user.RightsLevel != "" {
		_, err := db.Exec("UPDATE users SET rights_level = $1 WHERE id = $2", user.RightsLevel, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("Update user", user.ID)
	w.WriteHeader(http.StatusNoContent)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	// может быть, вынести в отдельную функцию? <
	userID := mux.Vars(r)["userid"]
	id_from_token, err := checkToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id_from_request, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if id_from_token != id_from_request {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	// >

	var user User
	err = db.QueryRow("SELECT id, login FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = db.QueryRow("SELECT name, profile_pic_path FROM employees WHERE id = $1", userID).Scan(&user.Name, &user.ProfilePicPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("Get user", userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func deleteUserByID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userid"]
	id_from_token, err := checkToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id_from_request, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if id_from_token != id_from_request {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Delete user", userID)
	w.WriteHeader(http.StatusNoContent)
}
