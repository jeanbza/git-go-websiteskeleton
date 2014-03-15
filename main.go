package main

import (
    "github.com/gorilla/mux"
    "git-go-websiteskeleton/app/home"
    "git-go-websiteskeleton/app/user"
    "net/http"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", home.GetHomePage).Methods("GET")
    r.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

    r.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    r.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}