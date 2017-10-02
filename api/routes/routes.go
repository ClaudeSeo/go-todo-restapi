package routes

import (
  "github.com/go-martini/martini"
  "github.com/claudeseo/go-todo-restapi/api/handler"
)

func URLPatterns (r martini.Router) {
    r.Get("/projects", handler.GetProjects)
    r.Post("/projects", handler.CreateProject)
    r.Get("/projects/(?P<id>[0-9]+)", handler.GetProject)
    r.Delete("/projects/(?P<id>[0-9]+)", handler.DeleteProject)
    r.Put("/projects/(?P<id>[0-9]+)/archive", handler.ArchiveProject)
    r.Put("/projects/(?P<id>[0-9]+)/restore", handler.RestoreProject)

    r.Get("/projects/(?P<projectId>[0-9]+)/tasks", handler.GetTasks)
    r.Post("/projects/(?P<projectId>[0-9]+)/tasks", handler.CreateTask)
    r.Get("/projects/(?P<projectId>[0-9]+)/tasks/(?P<taskId>[0-9]+)", handler.GetTask)
    r.Delete("/projects/(?P<projectId>[0-9]+)/tasks/(?P<taskId>[0-9]+)", handler.DeleteTask)
    r.Put("/projects/(?P<projectId>[0-9]+)/tasks/(?P<taskId>[0-9]+)/complete", handler.CompleteTask)
    r.Put("/projects/(?P<projectId>[0-9]+)/tasks/(?P<taskId>[0-9]+)/undo", handler.UndoTask)
}