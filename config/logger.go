package config

import (
	"log"
	"os"
)

const (
	logPath = "log.txt"
)

var Logger = log.New(os.Stdout, "", log.LstdFlags)
