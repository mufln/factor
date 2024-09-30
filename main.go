package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func initDB() {
	var err error
	connStr := "user=postgres password=postgres dbname=factor sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error during connecting db: ", err.Error())
		panic(err)
	}
}

func checkToken(r *http.Request) (int, error) {
	h, err := r.Cookie("Bearer")
	if err != nil {
		fmt.Println("No bearer in cookie,", err.Error())
		return 0, err
	}

	t := h.Value
	var tc TokenClaims
	token, err := jwt.ParseWithClaims(t, &tc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Unable to parse token,", err.Error())
			return nil, errors.New("Invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		fmt.Println("Unable to parse token,", err.Error())
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		fmt.Println("token claims are not of type *TokenClaims")
		return 0, errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserID, nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Это Factor. Тут пока ничего нет."))
	fmt.Println("/", r.Method)
}

func startHanding() {
	router := mux.NewRouter()
	router.HandleFunc("/", greetings).Methods("GET")
	router.HandleFunc("/users/", getUsers).Methods("GET")
	router.HandleFunc("/users/", createUser).Methods("POST")
	router.HandleFunc("/users/", updateUser).Methods("PUT")
	router.HandleFunc("/users/{userid}/", getUserByID).Methods("GET")
	router.HandleFunc("/users/{userid}/", deleteUserByID).Methods("DELETE")
	router.HandleFunc("/employees/", getEmployees).Methods("GET")
	router.HandleFunc("/employees/{userid}/", getEmployeeByID).Methods("GET")
	router.HandleFunc("/employees/{userid}/", updateEmployee).Methods("PUT")
	router.HandleFunc("/users/{userid}/chats/{GroupID}/{interval}/", getMessages).Methods("GET")
	router.HandleFunc("/users/{userid}/chats/{GroupID}/", sendMessage).Methods("POST")
	router.HandleFunc("/users/{userid}/chats/", getChats).Methods("GET")
	router.HandleFunc("/groups/", getGroups).Methods("GET")
	router.HandleFunc("/groups/", createGroup).Methods("POST")
	router.HandleFunc("/groups/{groupid}/", getGroupByID).Methods("GET")
	router.HandleFunc("/groups/{groupid}/", updateGroup).Methods("PUT")
	router.HandleFunc("/groups/{groupid}/", deleteGroup).Methods("DELETE")
	router.HandleFunc("/groups/{groupid}/members/", getGroupmembers).Methods("GET")
	router.HandleFunc("/groups/{groupid}/members/{userid}/", updateGroupmember).Methods("PUT")
	router.HandleFunc("/groups/{groupid}/members/{userid}/", deleteGroupmember).Methods("DELETE")
	router.HandleFunc("/users/{userid}/tasks/{from_date}/{to_date}/", getTasksByUserID).Methods("GET")
	router.HandleFunc("/users/{userid}/tasks/", editTask).Methods("PUT")
	router.HandleFunc("/users/{userid}/tasks/", deleteTask).Methods("DELETE")
	router.HandleFunc("/invites/{link}/", checkInvite).Methods("GET")
	router.HandleFunc("/register/{link}/", register).Methods("PUT")
	router.HandleFunc("/login/", login).Methods("POST")
	//router.HandleFunc("/logout/", logout).Methods("POST")
	//router.HandleFunc("/users/{userid}/tasks/{from_date}/{to_date}", getTasksByUserID).Methods("GET")
	//router.HandleFunc("/users/{userid}/tasks/", createTask).Methods("PUT")
	http.Handle("/", router)
}

func main() {
	initDB()
	startHanding()
	fmt.Println("Running on http://localhost:5000/")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err.Error())
	}
}
