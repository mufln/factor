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
	rows_users_table, err1 := db.Query("SELECT id, login FROM users")
	rows_employees_table, err2 := db.Query("SELECT name, profile_pic_path FROM employees")
	if err1 != nil || err2 != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows_users_table.Close()
	defer rows_employees_table.Close()

	for rows_users_table.Next() {
		rows_employees_table.Next()
		var user User
		if err := rows_users_table.Scan(&user.ID, &user.Login); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := rows_employees_table.Scan(&user.Name, &user.ProfilePicPath); err != nil {
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
