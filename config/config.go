package config

import (
	"time"
)

// Email unit struct
type EmailModel struct {
	Agent    string
	User     string
	Password string
	Host     string
	Port     string
}

// Server config
type ServerModel struct {
	Domain string
	Port   string
}

// Sender list
var (
	EmailConfig = []EmailModel{
		{"163", "go_sender@163.com", "tudyzhb", "smtp.163.com", "25"},
	}

	ServerConfig = ServerModel{"localhost", "6666"}
	// Chanel list
	// TaskChs = make([]chan EmailInfo, 10)
	SendEmailTimeOut = 20 * time.Second
)

// Parser config file
func ParserConfigFile(f_p string) (err error) {
	return
}
