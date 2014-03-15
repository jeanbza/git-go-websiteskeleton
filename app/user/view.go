package user

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
)

func GetViewPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title  string
        UserId string
    }

    params := mux.Vars(req)
    userId := params["id"]

    p := Page{
        Title:  "user_view",
        UserId: userId,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["user_view.html"] = template.Must(template.ParseFiles("templates/user/view.html", "templates/layout.html"))
    tmpl["user_view.html"].ExecuteTemplate(rw, "base", p)
}