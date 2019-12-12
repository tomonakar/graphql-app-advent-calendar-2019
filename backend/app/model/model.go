package model

import (
	"time"

	v "gopkg.in/go-playground/validator.v9"
)

var validator *v.Validate

func init() {
	validator = v.New()
}

type Task struct {
	ID         int64
	Identifier string
	Title      string
	Notes      string
	Completed  bool
	Due        time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (Task) IsNode() {}

func (t *Task) BeforeSave() error {
	return validator.Struct(t)
}
