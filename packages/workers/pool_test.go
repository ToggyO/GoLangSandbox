package workers

import (
    "hello/packages/workers/models"
    workers "hello/packages/workers/pool"
)

var add models.TaskFunc = func(a, b int) int {
    return a + b
}

func Run() {
    task := models.NewTask(func() {

    })

    pool := workers.NewPool(5)
    pool.Run()
}
