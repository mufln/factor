package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []User
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var emp User
		if err := rows.Scan(&emp.ID, &emp.ProfilePicPath, &emp.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	fmt.Println("Get employees")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
	w.WriteHeader(http.StatusOK)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
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

	var emp User
	rows_groups, err := db.Query("SELECT group_id FROM groupmembers WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer rows_groups.Close()

	for rows_groups.Next() {
		var group_id int
		if err := rows_groups.Scan(&group_id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		emp.Groups = append(emp.Groups, group_id)
	}

	err = db.QueryRow("SELECT rights_level FROM users WHERE id = $1", userID).Scan(&emp.RightsLevel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("Get employee", userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
	w.WriteHeader(http.StatusOK)
}

// Может быть, объединить эту функцию с updateUser?
func updateEmployee(w http.ResponseWriter, r *http.Request) {
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

	var emp User
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if emp.RightsLevel != "" {
		_, err := db.Exec("UPDATE users SET rights_level = $1 WHERE id = $2", emp.RightsLevel, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if len(emp.Groups) != 0 {
		for i := range emp.Groups {
			_, err := db.Exec("UPDATE groupmembers SET group_id = $1 WHERE id = $2 AND NOT EXISTS (SELECT * FROM groupmembers WHERE group_id = $1, user_id = $2)", i, userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	fmt.Println("Update employee", userID)
	w.WriteHeader(http.StatusNoContent)
}
