package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getGroupmembers(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["groupid"]
	if _, err := strconv.Atoi(groupID); err != nil {
		fmt.Println("Group ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	members := []Groupmember{}
	rows, err := db.Query("SELECT * FROM groupmembers WHERE group_id = $1", groupID)
	if err != nil {
		fmt.Println("Error during selecting from DB,", err.Error())
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var member Groupmember
		var role sql.NullString
		var g int
		err = rows.Scan(&g, &member.User.ID, &role)
		if err != nil {
			fmt.Println("Error during scanning data", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if role.Valid {
			member.Role = role.String
		}

		var u [4]sql.NullString
		err = db.QueryRow("SELECT u.id,u.login,u.rights_level,e.name,e.profile_pic_path FROM users AS u INNER JOIN employees AS e ON u.id = e.id AND u.id = $1;", member.User.ID).Scan(&member.User.ID, &u[0], &u[1], &u[2], &u[3])
		if err != nil {
			fmt.Println("Error during scanning data", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if u[0].Valid {
			member.User.Login = u[0].String
		}
		if u[1].Valid {
			member.User.RightsLevel = u[1].String
		}
		if u[2].Valid {
			member.User.Name = u[2].String
		}
		if u[3].Valid {
			member.User.ProfilePicPath = u[3].String
		}
		members = append(members, member)
	}

	fmt.Println("Got groupmembers")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	w.WriteHeader(http.StatusOK)
}

func updateGroupmember(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group_id := mux.Vars(r)["groupid"]
	groupID, err := strconv.Atoi(group_id)
	if err != nil {
		fmt.Println("Group ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user_id = mux.Vars(r)["userid"]
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println("User ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var member Groupmember
	if err = json.NewDecoder(r.Body).Decode(&member); err != nil {
		fmt.Println("Unable to decode json to Group struct,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var s int
	err = db.QueryRow("SELECT user_id FROM groupmembers WHERE group_id = $1", groupID).Scan(&s)
	if err != nil {
		fmt.Println("Unable to check if groupmember exists,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s != 0 {
		if member.Role != "" {
			_, err = db.Exec("UPDATE groupmembers SET role = $1 WHERE group_id = $2 AND user_id = $3", member.Role, groupID, userID)
			if err != nil {
				fmt.Println("Error during updating group_pic,", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		fmt.Println("No such groupmember")
		http.Error(w, "No such groupmember", http.StatusNotFound)
		return
	}

	fmt.Println("User", userID, "from group", groupID, "updated")
	w.WriteHeader(http.StatusOK)

}

func deleteGroupmember(w http.ResponseWriter, r *http.Request) {
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user_id string
	err = db.QueryRow("SELECT id FROM users WHERE id = $1", id_from_token).Scan(&user_id)
	if err != nil {
		fmt.Println("User does not exist,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group_id := mux.Vars(r)["groupid"]
	groupID, err := strconv.Atoi(group_id)
	if err != nil {
		fmt.Println("Group ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user_id = mux.Vars(r)["userid"]
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println("User ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = db.Exec("DELETE FROM groupmembers WHERE group_id = $1 AND user_id = $2", groupID, userID)
	if err != nil {
		fmt.Println("Unable to delete groupmember,", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Deleted groupmember", userID, "from group", groupID)
	w.WriteHeader(http.StatusOK)
}
