package workers

type task struct {
	Job        job
	Params     []param
	ResultChan chan interface{}
}

func newTask(job job, params []interface{}, resultChan chan interface{}) *task {
	l := len(params)
	paramValues := make([]param, l)

	for i := 0; i < l; i++ {
		paramValues[i] = newParamFromInterface(params[i])
	}

	return &task{job, paramValues, resultChan}
}
