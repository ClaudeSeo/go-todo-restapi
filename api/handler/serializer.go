package handler

import (
    "github.com/claudeseo/go-todo-restapi/api/database"
)

type ProjectSchma struct {
    ID uint `json:"id"`
    Title string `json:"title,omitempty"`
    Archive *bool `json:"archive,omitempty"`
    CreatedAt int64 `json:"create_at,omitempty"`
}

type TaskSchema struct {
    ID uint `json:"id"`
    Title string `json:"title,omitempty"`
    Done *bool `json:"done,omitempty"`
    CreatedAt int64 `json:"create_at,omitempty"`
    Deadline int64 `json:"dead_line,omitempty"`
}

type errorSchema struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Reason string `json:"reason,omitempty"`
}

type errorCode struct {
    Common map[string]errorSchema 
    Project map[string]errorSchema
    Task map[string]errorSchema
}

var ErrorCode errorCode

func (err errorSchema) addReason(reason string) errorSchema {
    err.Reason = reason
    return err
}

func MarshalProject (project database.Project, depth int) ProjectSchma {
    if depth <= 1 {
        return ProjectSchma{ID: project.ID}
    }
    archive := new(bool)
    *archive = project.Archive
    return ProjectSchma{project.ID, project.Title, archive, project.CreatedAt.Unix()}
}

func MarshalTask (task database.Task, depth int) TaskSchema {
    if depth <= 1 {
        return TaskSchema{ID: task.ID}
    }
    done := new(bool)
    *done = task.Done
    return TaskSchema{task.ID, task.Title, done, task.CreatedAt.Unix(), task.Deadline.Unix()}
}

func init () {
    ErrorCode.Common = map[string]errorSchema {
        "Unknown": errorSchema{Code: 10001, Message: "알 수 없는 에러가 발생했습니다."},
        "InvalidParameter": errorSchema{Code: 10002, Message: "잘못된 파라미터 입니다."},
    }
    ErrorCode.Project = map[string]errorSchema {
        "Unknown": errorSchema{Code: 20001, Message: "알 수 없는 프로젝트입니다."},
        "RequiredTitle": errorSchema{Code: 20002, Message: "타이틀은 필수입니다."},
    }
    ErrorCode.Task = map[string]errorSchema {
        "Unknown": errorSchema{Code: 30001, Message: "알 수 없는 작업입니다."},
        "RequiredTitle": errorSchema{Code: 30002, Message: "타이틀은 필수입니다."},
        "RequiredDeadline": errorSchema{Code: 30003, Message: "데드라인은 필수입니다."},
    }
}