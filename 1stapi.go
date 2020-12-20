package main

import (
    "fmt"
    "log"
    "net/http"
)

type User struct {
    id int `json:"id"`
    Name string `json:"Name"`
    Date_of_birth string `json:"Date_of_Birth`
    Phone_No. string `json:"Phone_No."`
    Email_Add string `json:"Email_Add"`
    Creation_Time int `json:"Creation_Time"`
}

var Users []User

type Contact struct {
    userOne User struct 
    userTwo User struct 
    var Timestamp int
}

func handleRequests() {
    http.HandleFunc("/", homePage)

    http.HandleFunc("/articles", returnAllArticles)
    log.Fatal(http.ListenAndServe(":27017", nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}


func main() {
    Users = []User{
        User1{id: 1, Name: "Sam", Date_of_Birth: "13/02/1987"},
    }
    handleRequests()
}

func main() {
    handleRequests()
}