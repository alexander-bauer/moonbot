package main

import (
	_ "encoding/json"
	"log"
	"os"
)

var (
	Version    = "0.2.1"
	minversion string
	l          *log.Logger
	quitSig    = make(chan int)

	bot *MoonBot

	ConfigFile string
)

const (
	defaultConfigFile = "moonbot.conf"
)

func main() {
	l = log.New(os.Stdout, "", log.Ltime)
	l.Println("MoonBot version:", Version+minversion)

	var err error
	bot, err = NewBot(os.Args[1], os.Args[2], l, "#moonbot")
	if err != nil {
		l.Fatalln("Could not create bot:", err)
	}
	l.Println("Bot created.") //This should list channels.

	status := <-quitSig
	os.Exit(status)
}
