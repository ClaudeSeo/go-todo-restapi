package main

import (
    "github.com/go-martini/martini"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

func main() {
    m := martini.Classic()
    db := database.InitDB()
    m.Map(db)
    m.Run()
}
