package v1

import "time"

func RunProgressBar() {
	var bar Bar
	bar.NewOption(0, 1200)
	//for i := 0; i < 3; i++ {
	//	time.Sleep(1000 * time.Millisecond)
	//	bar.Play(int64(400))
	//}
	bar.Play(int64(400))
	time.Sleep(1000 * time.Millisecond)

	bar.Play(int64(250))
	time.Sleep(1000 * time.Millisecond)

	bar.Play(int64(356))
	time.Sleep(1000 * time.Millisecond)

	bar.Play(int64(144))
	time.Sleep(1000 * time.Millisecond)

	bar.Play(int64(50))
	time.Sleep(1000 * time.Millisecond)

	bar.Finish()
	//var bar Bar
	//bar.NewOption(0, 100)
	//for i := 0; i <= 100; i++ {
	//	time.Sleep(100 * time.Millisecond)
	//	bar.Play(int64(i))
	//}
	//bar.Finish()
}
