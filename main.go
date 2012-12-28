package main

import (
	_ "encoding/json"
	"log"
	"os"
)

var (
	Version    = "0.1"
	minversion string
	l          *log.Logger
	quitSig    = make(chan int)

	ConfigFile string
)

const (
	defaultConfigFile = "moonbot.conf"
)

func main() {
	l = log.New(os.Stdout, "", log.Ltime)
	l.Println("MoonBot version:", Version+minversion)

	bot, err := NewBot(os.Args[1], os.Args[2], "#moonbot")
	if err != nil {
		l.Fatalln("Could not create bot:", err)
	}
	l.Println("Bot created.") //This should list channels.

	l.Println("Saying hello.")
	bot.Say("#moonbot", "Hello.")

	status := <-quitSig
	os.Exit(status)
}
