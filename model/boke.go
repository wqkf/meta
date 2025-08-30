package main

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	UserName string
	Password string
	email    string
	Posts    []Post
	Comments []Comment
	gorm.Model
}

type Post struct {
	ID       uint
	Title    string
	Content  string
	UserID   uint
	Comments []Comment
	gorm.Model
}

type Comment struct {
	ID      uint
	PostId  uint
	UserID  uint
	Content string
	gorm.Model
}
