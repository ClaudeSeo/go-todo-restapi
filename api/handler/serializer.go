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

type errorSchema struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Reason string `json:"reason"`
}

type errorCode struct {
    Common map[string]errorSchema 
    Project map[string]errorSchema
}

var ErrorCode errorCode

func (err errorSchema) addReason(reason string) errorSchema {
    err.Reason = reason
    return err
}

func MarshalProject (project database.Project, depth int) ProjectSchma {
    archive := new(bool)
    *archive = project.Archive
    if depth <= 1 {
        return ProjectSchma{ID: project.ID}
    }
    return ProjectSchma{project.ID, project.Title, archive, project.CreatedAt.Unix()}
}


func init () {
    ErrorCode.Common = map[string]errorSchema {
        "Unknown": errorSchema{Code: 10001, Message: "알 수 없는 에러가 발생했습니다."},
        "InvalidParameter": errorSchema{Code: 10002, Message: "잘못된 파라미터 입니다."},
    }
    ErrorCode.Project = map[string]errorSchema {
        "RequiredTitle": errorSchema{Code: 20001, Message: "타이틀은 필수입니다."},
    }
}