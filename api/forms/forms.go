package forms

import "github.com/martini-contrib/binding"

type ProjectForm struct {
    Title string `json:"title" binding:"required"`
}
