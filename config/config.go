package config

// Email unit struct
type EmailModel struct {
	Agent    string
	User     string
	Password string
	Host     string
	Port     string
}

// Sender list
var (
	EmailConfig = []EmailModel{
		{"163", "go_sender@163.com", "tudyzhb", "smtp.163.com", "25"},
	}
)
