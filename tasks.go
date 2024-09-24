package main

//func getTasksByUserID(w http.ResponseWriter, r *http.Request) {
//	//	/users/{userid}/tasks/{from_date}/{to_date}
//	userID := mux.Vars(r)["userid"]
//	fromDate := mux.Vars(r)["from_date"]
//	toDate := mux.Vars(r)["to_date"]
//
//	var tasks []Task
//
//	rows, err := db.Query("SELECT id, from_user, state, task_text, created_at, completed_at FROM tasks WHERE to_user = $1", userID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var task Task
//		var created_at int
//		if err := rows.Scan(&task.ID, &task.FromUserID, &task.IsCompleted, &task.Text, &created_at, &task.CompletedAt); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		if db.QueryRow("SELECT name FROM employees WHERE id = $1", task.FromUserID).Scan(&task.FromUserName) != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		tasks = append(tasks, task)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(tasks)
//}
