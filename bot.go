package main

import (
	"github.com/zeebo/irc"
)

type MoonBot struct {
	Conn *irc.Connection
}

const (
	Nick    = "MoonBot"
	AltNick = "MoonBawt"
)

func NewBot(channel, server, password string) (bot *MoonBot, err error) {
	info := irc.Info{
		Channel:  channel,
		Nick:     Nick,
		AltNick:  AltNick,
		Server:   server,
		Password: password,
	}

	ircconn, err := irc.NewConnection(info)
	if err != nil {
		return nil, err
	}

	bot = &MoonBot{
		Conn: ircconn,
	}
	return
}
