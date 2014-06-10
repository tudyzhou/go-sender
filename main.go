package main

import (
	"github.com/tudyzhou/go-sender/config"
	"github.com/tudyzhou/go-sender/sender"
	"net/http"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) == 1 {
		server()
	} else {
		api := sender.Sender
		info := &sender.EmailInfo{"go_sender@163.com", "中文标题", "这是中文", ""}
		err := api.Send(info)
		if err != nil {
			println(err)
		}
	}
}

// http listen server
func server() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	addr := config.ServerConfig.Domain + ":" + config.ServerConfig.Port
	http.HandleFunc("/", sender.TaskHandler)

	config.Logger.Println("[Debug] SendMail server listen at", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		config.Logger.Println("[ERROR]", err)
	} else {
		config.Logger.Panicln("[Debug] SendMail server listen at", addr)
	}
}

// background server
func task() {
	println("aaa")
}
