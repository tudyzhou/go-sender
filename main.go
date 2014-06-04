package main

import (
	"github.com/tudyzhou/go-sender/sender"
)

func main() {
	api := sender.Sender
	info := &sender.EmailInfo{"go_sender@163.com", "中文标题", "这是中文", ""}
	err := api.Send(info)
	if err != nil {
		println(err)
	}
}
