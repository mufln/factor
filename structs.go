package main

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
)

var db *sql.DB

const signingKey = "testkey"

type User struct {
	ID             int    `json:"id" db:"id"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	RightsLevel    string `json:"rights_level"`
	ProfilePicPath string `json:"profile_pic_path"`
	Name           string `json:"name"`
	Groups         []int  `json:"groups"`
}

type Message struct {
	FromUserID  int    `json:"from_user_id"`
	GroupID     int    `json:"group_id"`
	MessageText string `json:"message_text"`
	CreatedAt   int    `json:"created_at"`
}

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	GroupPic string `json:"group_pic"`
}

type Task struct {
	ID           int    `json:"id"`
	IsCompleted  bool   `json:"is_completed"`
	IsOverdue    bool   `json:"is_overdue"`
	CompletedAt  int    `json:"completed_at"`
	Text         string `json:"text"`
	FromUserID   int    `json:"from_user_id"`
	FromUserName string `json:"from_user_name"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"id"`
}

type Token struct {
	SignedToken string `json:"signed_token"`
}

//type Link struct {
//	linkExists bool `json:"link_exists"`
//}
