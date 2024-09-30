package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func getTasksByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userid"]
	fromDate := mux.Vars(r)["from_date"]
	toDate := mux.Vars(r)["to_date"]

	//id_from_token, err := checkToken(r)
	_, err := checkToken(r)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//id_from_request, err := strconv.Atoi(userID)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//if id_from_token != id_from_request {
	//	http.Error(w, err.Error(), http.StatusUnauthorized)
	//	return
	//}

	var tasks []Task

	rows, err := db.Query("SELECT id, from_user, state, task_text, created_at, expired_at, completed_at FROM tasks WHERE to_user = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		//var created_at, expited_at, comleted_at time.Time
		//var text string
		//var state bool
		//  state        | boolean                     |           |          |
		//	task_text    | text                        |           |          |
		//	created_at   | timestamp without time zone |           |          |
		//	expired_at   | timestamp without time zone |           |          |
		//	completed_at | timestamp without time zone |           |          |

		var info [5]sql.NullString
		err = rows.Scan(&task.ID, &task.FromUserID, &task.ToUserId, &info[0], &info[1], &info[2], &info[3], &info[4])
		if err != nil {
			fmt.Println("Error during connecting DB", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// &created_at, &expited_at, &comleted_at

		if info[0].Valid {
			task.IsCompleted, err = strconv.ParseBool(info[0].String)
			if err != nil {
				fmt.Println("Error during converting state", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if info[1].Valid {
			task.Text = info[1].String
		}
		if info[2].Valid {
			// parse timestamp to unix time
			// i'm too tired for today, sorry :(
		}
		//if created_at.Unix() != 0{
		//	task.CreatedAt = int(created_at.Unix())
		//}
		//if comleted_at.Unix() != 0{
		//	task.CompletedAt = int(comleted_at.Unix())
		//}
		//if expited_at.Unix() != 0{
		//	task.ExpiredAt = int(expited_at.Unix())
		//}

		if task.ExpiredAt >= task.CompletedAt {
			task.IsOverdue = false
		} else {
			task.IsOverdue = true
		}

		if db.QueryRow("SELECT name FROM employees WHERE id = $1", task.FromUserID).Scan(&task.FromUserName) != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		from, err := strconv.Atoi(fromDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		to, err := strconv.Atoi(toDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if task.CreatedAt >= from && task.CreatedAt <= to {
			tasks = append(tasks, task)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func editTask(w http.ResponseWriter, r *http.Request) {
	user_id := mux.Vars(r)["userid"]
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println("User ID incorrect,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id_from_token, err := checkToken(r)
	fmt.Println("ID from token:", id_from_token)
	if err != nil {
		fmt.Println("Token invalid,", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// если пользователь может просматривать и изменять только свои задачи,
	// то сюда нужно добавить проверку на соответствие {userid} и id из токена,

	err = db.QueryRow("SELECT id FROM users WHERE id = $1", userID).Scan(userID)
	if err != nil {
		fmt.Println("Error during connecting DB", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if userID == 0 {
		fmt.Println("No such user")
		http.Error(w, "No user with provided {userid}", http.StatusNotFound)
		return
	}

	var task Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println("Unable to decode request,", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var doesExist string
	var isCompleted bool
	err = db.QueryRow("SELECT to_user, state FROM tasks WHERE id = $1;", task.ID).Scan(&doesExist, &isCompleted)
	if err != nil {
		fmt.Println("Error during connecting DB", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if doesExist == "" {
		fmt.Println("Task does not exist, creating")
		if task.FromUserID == 0 {
			fmt.Println("No from_user_id provided, it will be replaced with {userid}")
			task.FromUserID = userID
		}
		if task.ToUserId == 0 {
			fmt.Println("No to_user_id provided, it will be replaced with {userid}")
			task.ToUserId = userID
		}
		_, err = db.Exec("INSERT INTO tasks (from_user, to_user, created_at) VALUES ($1, $2, $3)", task.FromUserID, task.ToUserId, time.Now().Unix())
		if err != nil {
			fmt.Println("Unable to create task,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.IsCompleted != isCompleted {
		_, err = db.Exec("UPDATE tasks SET state = $1 WHERE id = $2", task.IsCompleted, task.ID)
		if err != nil {
			fmt.Println("Unable to update is_completed field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.ExpiredAt != 0 {
		_, err = db.Exec("UPDATE tasks SET expired_at = $1 WHERE id = $2", task.ExpiredAt, task.ID)
		if err != nil {
			fmt.Println("Unable to update expired_at field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.CompletedAt != 0 {
		_, err = db.Exec("UPDATE tasks SET completed_at = $1 WHERE id = $2", task.CompletedAt, task.ID)
		if err != nil {
			fmt.Println("Unable to update completed_at field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.Text != "" {
		_, err = db.Exec("UPDATE tasks SET text = $1 WHERE id = $2", task.Text, task.ID)
		if err != nil {
			fmt.Println("Unable to update text field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.FromUserID != 0 {
		_, err = db.Exec("UPDATE tasks SET from_user = $1 WHERE id = $2", task.FromUserID, task.ID)
		if err != nil {
			fmt.Println("Unable to update from_user field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if task.ToUserId != 0 {
		_, err = db.Exec("UPDATE tasks SET to_user = $1 WHERE id = $2", task.ToUserId, task.ID)
		if err != nil {
			fmt.Println("Unable to update to_user field,", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if doesExist == "" {
		fmt.Println("task created")
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Println("task updated")
		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}
