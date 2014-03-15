package home

import (
    "net/http"
    "html/template"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "home",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["home.html"] = template.Must(template.ParseFiles("templates/home/home.html", "templates/layout.html"))
    tmpl["home.html"].ExecuteTemplate(rw, "base", p)
}