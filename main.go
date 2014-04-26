package main

import (
    "net/http"
    "time"

    "git-go-websiteskeleton/app/common"
    "git-go-websiteskeleton/app/home"
    "git-go-websiteskeleton/app/user"

    "github.com/gorilla/mux"
)

var router *mux.Router

const (
    accessLogging = true
)

func main() {
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

type ApacheLogRecord struct {
    http.ResponseWriter
 
    ip                    string
    time                  time.Time
    method, uri, protocol string
    status                int
    responseBytes         int64
    elapsedTime           time.Duration
}

func httpInterceptor(w http.ResponseWriter, req *http.Request) {
    startTime := time.Now()

    router.ServeHTTP(w, req)

    finishTime := time.Now()
    elapsedTime := finishTime.Sub(startTime)

    if (accessLogging) {
        common.LogAccess(w, req, elapsedTime)
    }
}