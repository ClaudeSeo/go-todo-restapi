package main

import (
    "fmt"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/claudeseo/go-todo-restapi/config"
    "github.com/claudeseo/go-todo-restapi/api/routes"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

func main() {
    m := martini.Classic()
    db := database.InitDB()
    m.Use(render.Renderer())    
    m.Map(db)
    m.Group("/api", routes.URLPatterns)
    m.RunOnAddr(fmt.Sprintf("%s:%d", config.WebConfig.Host, config.WebConfig.Port))
}
