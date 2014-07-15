package user

import (
    "net/http"
    "html/template"
    
    "git-go-websiteskeleton/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "user_home",
    }

    common.Templates = template.Must(template.ParseFiles("templates/user/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}