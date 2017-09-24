package routes

import (
  "github.com/go-martini/martini"
)

func URLPatterns (r martini.Router) {
    r.Get("/projects", "")
}