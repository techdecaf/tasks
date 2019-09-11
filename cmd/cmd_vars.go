package cmd

import (
	"github.com/techdecaf/golog"
	"github.com/techdecaf/tasks/internal/taskfile"
)

var log = golog.Log{
	Name: "tasks",
}

var tasks = &taskfile.TaskFile{}
