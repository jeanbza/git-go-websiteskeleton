package user

import (
    "net/http"
    "html/template"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "user_home",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["user_home.html"] = template.Must(template.ParseFiles("static/html/user/home.html", "static/html/layout.html"))
    tmpl["user_home.html"].ExecuteTemplate(rw, "base", p)
}