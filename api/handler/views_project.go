package handler

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/martini-contrib/render"
    "github.com/claudeseo/go-todo-restapi/api/forms"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

func GetProjects(r render.Render, req *http.Request, db *gorm.DB) {
    projects := []database.Project{}
    db.Order("id desc").Find(&projects)
    MarshalProjects(r, 200, projects)
}

func CreateProject(res http.ResponseWriter, req *http.Request, form forms.ProjectForm) {
    fmt.Println(form)
}