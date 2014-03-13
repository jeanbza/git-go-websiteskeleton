package home

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "home",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["home.html"] = template.Must(template.ParseFiles("static/html/home.html", "static/html/index.html"))
    tmpl["home.html"].ExecuteTemplate(rw, "base", p)
}