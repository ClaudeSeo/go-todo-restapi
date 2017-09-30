package routes

import (
  "github.com/go-martini/martini"
  "github.com/claudeseo/go-todo-restapi/api/handler"
)

func URLPatterns (r martini.Router) {
    r.Get("/projects", handler.GetProjects)
}