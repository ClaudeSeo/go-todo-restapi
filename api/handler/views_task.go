package handler

import (
    "time"
    "net/http"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/claudeseo/go-todo-restapi/api/database"
)

type TaskForm struct {
    Title string
    Deadline int64
}

func GetTasks(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    tasks := []database.Task{}
    if err := db.Model(&project).Related(&tasks).Error; err != nil {
        r.JSON(404, ErrorCode.Task["Unknown"])
    }

    var results []TaskSchema
    for _, task := range tasks {
        results = append(results, MarshalTask(task, 2))
    }
    r.JSON(200, map[string]interface{}{"data": results})
}

func CreateTask(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    form := TaskForm{}
    decoder := json.NewDecoder(req.Body)
    if err := decoder.Decode(&form); err != nil {
        r.JSON(400, ErrorCode.Common["InvalidParameter"].addReason(err.Error()))
        return
    }
    if form.Title == "" {
        r.JSON(400, ErrorCode.Task["RequiredTitle"])
        return
    }
    if form.Deadline == 0 {
        r.JSON(400, ErrorCode.Task["RequiredDeadline"])
        return
    }
    defer req.Body.Close()

    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    task := database.Task{Title: form.Title, Deadline: time.Unix(form.Deadline, 0), Done: false, ProjectId: project.ID}
    db.Create(&task)
    r.JSON(201, MarshalTask(task, 1))
}

func GetTask(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    task := database.Task{}
    if err := db.First(&task, map[string]interface{} {"id": params["taskId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Task["Unknown"])
        return
    }
    r.JSON(200, MarshalTask(task, 2))
}

func CompleteTask(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    task := database.Task{}
    if err := db.First(&task, map[string]interface{} {"id": params["taskId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Task["Unknown"])
        return
    }
    task.Done = true
    if err := db.Save(&task).Error; err != nil {
        r.JSON(500, ErrorCode.Common["Unknown"].addReason(err.Error()))
        return
    }
    r.JSON(200, MarshalTask(task, 2))
}

func UndoTask(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    task := database.Task{}
    if err := db.First(&task, map[string]interface{} {"id": params["taskId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Task["Unknown"])
        return
    }
    task.Done = false
    if err := db.Save(&task).Error; err != nil {
        r.JSON(500, ErrorCode.Common["Unknown"].addReason(err.Error()))
        return
    }
    r.JSON(200, MarshalTask(task, 2))
}

func DeleteTask(r render.Render, req *http.Request, db *gorm.DB, params martini.Params) {
    project := database.Project{}
    if err := db.First(&project, map[string]interface{} {"id": params["projectId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Project["Unknown"])
        return
    }
    task := database.Task{}
    if err := db.First(&task, map[string]interface{} {"id": params["taskId"]}).Error; err != nil {
        r.JSON(404, ErrorCode.Task["Unknown"])
        return
    }
    if err := db.Delete(&task).Error; err != nil {
        r.JSON(500, ErrorCode.Common["Unknown"].addReason(err.Error()))
        return
    }
    r.JSON(204, nil)
}
