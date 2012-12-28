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

	ConfigFile string
)

const (
	defaultConfigFile = "moonbot.conf"
)

func main() {
	l = log.New(os.Stdout, "", log.Ltime)
	l.Println("MoonBot version:", Version+minversion)

	_, err := NewBot("#moonbot", os.Args[1], os.Args[2])
	if err != nil {
		l.Fatalln("Could not create bot:", err)
	}
	l.Println("Bot created.") //This should list channels.

	l.Println("Successful.")
}
