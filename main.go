package main

import (
    "net/http"
    "time"
    "flag"

    "git-go-websiteskeleton/app/common"
    "git-go-websiteskeleton/app/home"
    "git-go-websiteskeleton/app/user"

    "github.com/gorilla/mux"
    "github.com/golang/glog"
)

var router *mux.Router

func main() {
    flag.Parse()
    defer glog.Flush()

    router = mux.NewRouter()
    http.HandleFunc("/", httpInterceptor)

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func httpInterceptor(w http.ResponseWriter, req *http.Request) {
    startTime := time.Now()

    router.ServeHTTP(w, req)

    finishTime := time.Now()
    elapsedTime := finishTime.Sub(startTime)

    switch req.Method {
    case "GET":
        // We may not always want to StatusOK, but for the sake of 
        // this example we will
        common.LogAccess(w, req, elapsedTime)
    case "POST":
        // here we might use http.StatusCreated
    }
}