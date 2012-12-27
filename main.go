package main

import (
	_ "encoding/json"
	_ "github.com/zeebo/irc"
)

var (
	Version     = "0.1"
	minveresion string

	ConfigFile string
)

const (
	defaultConfigFile = "moonbot.conf"
)

type MoonBot struct {
}

func main() {

}
