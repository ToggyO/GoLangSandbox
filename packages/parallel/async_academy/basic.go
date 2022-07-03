package async_academy

type (
	In   = chan<- int
	Out  = <-chan int
	Bi   = chan int
	Done = chan struct{}
)

func RunAcademyCode() {
	//RunAsyncWrite()
	RunAsyncRead()
}

func checkOnComplete(done Done) {
	complete := 0
	for _ = range done {
		complete++
		if complete == 3 {
			close(done)
		}
	}
}
