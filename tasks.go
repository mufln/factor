package main

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
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		to, err := strconv.Atoi(toDate)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		if task.CreatedAt >= from && task.CreatedAt <= to {
//			tasks = append(tasks, task)
//		}
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(tasks)
//}
//
//func createTask(w http.ResponseWriter, r *http.Request) {
//	userID := mux.Vars(r)["userid"]
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
//	var task Task
//	if err = json.NewDecoder(r.Body).Decode(&task); err != nil {
//		fmt.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	//{
//	//  "id": 128,
//	//  "is_completed": false,
//	//  "is_overdue": false,
//	//  "completed_at": 1083201980,
//	//  "text": "free durov",
//	//  "from_user_id": 0,
//	//  "from_user_name": "Pavel Durov"
//	//}
//	if task.ID != 0 {
//		err = db.Exec("UPDATE tasks SET id = $1, completed_at = $2, expired_at = $3")
//	}
//	//_, err = db.Exec("INSERT INTO tasks () VALUES ($1)", user.Login)
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusInternalServerError)
//	//	return
//	//}
//
//	fmt.Println("Create user")
//	w.WriteHeader(http.StatusCreated)
//}
