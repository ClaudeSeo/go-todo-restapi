package database

import (
    "fmt"
    "log"
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/claudeseo/go-todo-restapi/config"
)

type Project struct {
    gorm.Model
    Title string 
    Archive bool
    Tasks []Task `gorm:"ForeignKey:ProdjctId;AssociationForeignKey:Refer"`
}

type Task struct {
    gorm.Model
    Title string 
    Deadline time.Time
    Done bool
    ProjectId uint
}

func InitDB () *gorm.DB {
    uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
        config.DBConfig.Username,
        config.DBConfig.Password,
        config.DBConfig.Host,
        config.DBConfig.Port,
        config.DBConfig.Name,
    )
    db, err :=  gorm.Open(config.DBConfig.Dialect, uri)
    if err != nil {
        log.Fatal("Could not connect database")
        return nil
    }
    db = InitSchema(db)
    return db
}

func InitSchema (db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&Project{}, &Task{})
    db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
    return db
}