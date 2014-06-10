package sender

import (
	"net/http"
	"time"
)

var (
	chs = make([]chan EmailInfo, 10)
)

// parser request and send email
func TaskHandler(rw http.ResponseWriter, req *http.Request) {
	info := EmailInfo{"sender@163.com", "subject", "test body", ""}
	rw.WriteHeader(200)
	rw.Write([]byte("got it"))
	go sender_test(info)
}

func sender_test(info EmailInfo) {
	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	flg := make(chan bool, 1)
	go func() {
		println("test", info.Body)
		for i := 0; i < 1000000000000; i++ {
		} // Debug
		flg <- true
	}()

	select {
	case <-flg:
		println("finish")
		return
	case <-t.C:
		println("time out")
		return
	}
}
