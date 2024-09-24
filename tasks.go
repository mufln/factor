package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//func getTasksByUserID(w http.ResponseWriter, r *http.Request) {
//	//	/users/{userid}/tasks/{from_date}/{to_date}
//	userID := mux.Vars(r)["userid"]
//	fromDate := mux.Vars(r)["from_date"]
//	toDate := mux.Vars(r)["to_date"]
//	id_from_token, err := checkToken(r)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	id_from_request, err := strconv.Atoi(userID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	if id_from_token != id_from_request {
//		http.Error(w, err.Error(), http.StatusForbidden)
//		return
//	}
//
//	var tasks []Task
//
//	rows, err := db.Query("SELECT id, from_user, state, task_text, created_at, expired_at, completed_at FROM tasks WHERE to_user = $1", userID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var task Task
//		var created_at, expited_at, comleted_at time.Time
//		if err := rows.Scan(&task.ID, &task.FromUserID, &task.IsCompleted, &task.Text, &created_at, &expited_at, &comleted_at); err != nil {
//			fmt.Println(err.Error())
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		task.CreatedAt = int(created_at.Unix())
//		task.CompletedAt = int(comleted_at.Unix())
//		task.ExpiredAt = int(expited_at.Unix())
//
//		if task.ExpiredAt >= task.CompletedAt {
//			task.IsOverdue = false
//		} else {
//			task.IsOverdue = true
//		}
//
//		if db.QueryRow("SELECT name FROM employees WHERE id = $1", task.FromUserID).Scan(&task.FromUserName) != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		from, err := strconv.Atoi(fromDate)
//		if err != nil {
//
//		}
//		tasks = append(tasks, task)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(tasks)
//}
