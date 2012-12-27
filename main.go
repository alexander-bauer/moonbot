package main

import (
	_ "encoding/json"
	_ "github.com/zeebo/irc"
	"log"
	"os"
)

var (
	Version    = "0.1"
	minversion string

	ConfigFile string
)

const (
	defaultConfigFile = "moonbot.conf"
)

type MoonBot struct {
	l *log.Logger
}

func main() {
	bot := &MoonBot{
		l: log.New(os.Stdout, "moonbot:", log.Ltime),
	}
	bot.l.Println("MoonBot version:", Version+minversion)
}
