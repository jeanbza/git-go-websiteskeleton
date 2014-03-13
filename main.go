package main

import (
    "git-go-websiteskeleton/src/home"
    "git-go-websiteskeleton/src/common"
    "net/http"
)

func main() {
    http.HandleFunc("/", homePage)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    common.CheckError(err)
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}