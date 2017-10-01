package handler

import (
    "net/http"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/martini-contrib/render"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

type ProjectForm struct {
    Title string
}

func GetProjects(r render.Render, db *gorm.DB) {
    projects := []database.Project{}
    db.Order("id desc").Find(&projects)
    var results []ProjectSchma
    for _, project := range projects {
        results = append(results, MarshalProject(project, 2))
    }
    r.JSON(200, map[string]interface{}{"data": results})
}

func CreateProject(r render.Render, req *http.Request, db *gorm.DB) {
    form := ProjectForm{}
    decoder := json.NewDecoder(req.Body)
    if err := decoder.Decode(&form); err != nil {
        r.JSON(400, ErrorCode.Common["InvalidParameter"].addReason(err.Error()))
        return
    }
    if form.Title == "" {
        r.JSON(400, ErrorCode.Project["RequiredTitle"])
        return
    }
    defer req.Body.Close()
    project := database.Project{Title: form.Title}
    db.Create(&project)
    r.JSON(201, MarshalProject(project, 1))
}