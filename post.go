package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Post struct {
	gorm.Model
	Title  string
	Author string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Post{})
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var posts []Post
	db.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	db.Create(&Post{Title: vars["title"], Author: vars["author"]})

	fmt.Fprintf(w, "New Post Successfully Created")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)

	var post Post
	db.Where("title = ?", vars["title"]).Find(&post)
	db.Delete(&post)

	fmt.Fprintf(w, "Post Successfully Deleted")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)

	var post Post
	db.Where("ID = ?", vars["ID"]).Find(&post)

	post.Title = vars["title"]

	db.Save(&post)

	fmt.Fprintf(w, "Successfully Updated Post")
}
