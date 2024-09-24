package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func isUserInGroup(userID, GroupID string) bool {
	var v int
	err := db.QueryRow("SELECT group_id FROM groupmembers WHERE user_id = $1 AND group_id = $2", userID, GroupID).Scan(&v)
	if err != nil {
		return false
	}
	id, _ := strconv.Atoi(GroupID)
	if v != id {
		return false
	}
	return true
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userid"]
	groupID := mux.Vars(r)["GroupID"]
	interval := mux.Vars(r)["interval"]
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

	if !isUserInGroup(userID, groupID) {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	rows, err := db.Query("SELECT message_text, created_at, from_user_id FROM messages WHERE group_id = $1 ORDER BY created_at OFFSET $2 LIMIT 20", groupID, interval)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.MessageText, &message.CreatedAt, &message.FromUserID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(message)
		messages = append(messages, message)
	}

	fmt.Println("Get messages")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
	w.WriteHeader(http.StatusOK)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userid"]
	groupID := mux.Vars(r)["GroupID"]

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

	if !isUserInGroup(userID, groupID) {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = db.Exec("INSERT INTO messages (group_id, from_user_id, created_at, message_text) VALUES ($1, $2, $3, $4)", groupID, userID, message.CreatedAt, message.MessageText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Message wasn't sent,", err.Error())
		return
	}

	fmt.Println("Send message from", userID)
	w.WriteHeader(http.StatusCreated)
}

func getChats(w http.ResponseWriter, r *http.Request) {
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

	var groups []Group
	rows, err := db.Query("SELECT group_id FROM groupmembers WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var group Group
		if err := rows.Scan(&group.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.QueryRow("SELECT name, group_pic FROM groups WHERE id = $1", group.ID).Scan(&group.Name, &group.GroupPic)
		groups = append(groups, group)
	}

	fmt.Println("Get chats for", userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
	w.WriteHeader(http.StatusOK)
}
