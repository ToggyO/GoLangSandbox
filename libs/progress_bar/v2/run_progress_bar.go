package v2

import (
	"crypto/rand"
	"io"
	"io/ioutil"
)

func RunProgressBar() {
	var limit int64 = 1024 * 1024 * 500

	// we will copy 500 MiB from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar := NewProgressBar(limit)
	bar.Start()

	// create proxy reader
	barReader := bar.NewProxyReader(reader)

	// copy from proxy reader
	io.Copy(writer, barReader)

	// finish bar
	bar.Finish()
	//var limit int64 = 10000
	//
	//bar := NewProgressBar(limit)
	//bar.Start()
	//
	//for i := 0; i < int(limit); i++ {
	//	bar.Add(1)
	//	time.Sleep(time.Millisecond)
	//}
	//
	//bar.Finish()
}
