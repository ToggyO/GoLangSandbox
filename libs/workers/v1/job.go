package v1

type job struct {
	Function interface{}
}

func newJob(function interface{}) job {
	return job{function}
}
