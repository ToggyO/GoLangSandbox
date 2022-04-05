package models

type TaskFunc func()

//type TaskFunc func(args ...any) any

type Task struct {
    Execute TaskFunc
}

func NewTask(taskFunc TaskFunc) *Task {
    return &Task{taskFunc}
}
