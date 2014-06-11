package sender

import (
	"github.com/tudyzhou/go-sender/config"
	"net/http"
	"time"
)

// parser request and send email
func TaskHandler(rw http.ResponseWriter, req *http.Request) {
	info := EmailInfo{"go_sender@163.com", "subject", "test body", ""}
	rw.WriteHeader(200)
	rw.Write([]byte("got it"))
	go sender_go(&info)
}

func sender_go(info *EmailInfo) {
	// task time out
	t := time.NewTimer(config.SendEmailTimeOut)
	flg := make(chan bool, 1)
	defer t.Stop()

	go func() {
		//for i := 0; i < 1000000000000; i++ {
		//} // Debug
		err := Sender.Send(info)
		if err != nil {
			config.Logger.Println("[ERROR] send email error", err)
		} else {
			flg <- true
		}
	}()

	select {
	case <-flg:
		config.Logger.Println("[DETAIL] send email success")
		return
	case <-t.C:
		// todo: time out, and send again
		config.Logger.Println("[ERROR] time out")
		return
	}
}
