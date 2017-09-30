package handler

import (
    "time"
    "github.com/martini-contrib/render"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

type ProjectSchma struct {
    ID uint `json:"id"`
    Title string `json:"title"`
    Archive bool `json:"archive"`
    CreatedAt time.Time `json:"create_at"`
}

func MarshalProjects (r render.Render, status int, projects []database.Project) {
    var results []ProjectSchma
    for _, project := range projects {
        results = append(results, ProjectSchma{project.ID, project.Title, project.Archive, project.CreatedAt})
    }
    r.JSON(status, map[string]interface{}{"data": results})
}
