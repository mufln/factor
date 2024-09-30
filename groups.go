package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getGroups(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groups := []Group{}
	rows, err := db.Query("SELECT * FROM groups;")
	if err != nil {
		fmt.Println("Error during selecting from DB,", err.Error())
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var group Group
		var n, p sql.NullString
		if err = rows.Scan(&group.ID, &n, &p); err != nil {
			fmt.Println("Error during scanning data", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if n.Valid {
			group.Name = n.String
		}
		if p.Valid {
			group.GroupPic = p.String
		}
		groups = append(groups, group)
	}

	fmt.Println("Got groups")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
	w.WriteHeader(http.StatusOK)
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO groups (name) VALUES ('default name');")
	if err != nil {
		fmt.Println("Error during inserting in DB", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Created group")
	w.WriteHeader(http.StatusOK)
}

func getGroupByID(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID := mux.Vars(r)["groupid"]
	id, err := strconv.Atoi(groupID)
	if err != nil {
		fmt.Println("Unable to convert {groupid} to int,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var group Group
	var n, p sql.NullString
	err = db.QueryRow("SELECT * FROM groups WHERE id = $1;", id).Scan(&group.ID, &n, &p)
	if err != nil {
		fmt.Println("Error during selecting from DB,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if n.Valid {
		group.Name = n.String
	}
	if p.Valid {
		group.GroupPic = p.String
	}
	fmt.Println("Get group", groupID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
	w.WriteHeader(http.StatusOK)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var group Group
	groupID := mux.Vars(r)["groupid"]
	group.ID, err = strconv.Atoi(groupID)
	if err != nil {
		fmt.Println("Unable to convert {groupid} to int,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&group); err != nil {
		fmt.Println("Unable to decode json to Group struct,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var s int
	err = db.QueryRow("SELECT id FROM groups WHERE id = $1", group.ID).Scan(&s)
	if err != nil {
		fmt.Println("Unable to check if group exists,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s != 0 {
		if group.GroupPic != "" {
			_, err = db.Exec("UPDATE groups SET group_pic = $1 WHERE id = $2", group.GroupPic, group.ID)
			if err != nil {
				fmt.Println("Error during updating group_pic,", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if group.Name != "" {
			_, err = db.Exec("UPDATE groups SET name = $1 WHERE id = $2", group.Name, group.ID)
			if err != nil {
				fmt.Println("Error during updating name,", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		_, err = db.Exec("INSERT INTO groups (name, group_pic) VALUES ($1, $2)", group.Name, group.GroupPic)
		if err != nil {
			fmt.Println("Error during creating group,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Println("Group", group.ID, "updated")
	w.WriteHeader(http.StatusNoContent)
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID := mux.Vars(r)["groupid"]
	_, err = db.Exec("DELETE FROM groups WHERE id = $1", groupID)
	if err != nil {
		fmt.Println("Unable to delete group,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Deleted group", groupID)
	w.WriteHeader(http.StatusOK)
}
