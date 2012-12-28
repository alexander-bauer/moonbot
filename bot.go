package main

import (
	"github.com/museun/irc"
	"log"
)

type MoonBot struct {
	Conn *irc.Conn

	l *log.Logger
}

const (
	Nick = "MoonBot"
	Real = "Duo's Bot"
)

func NewBot(server, password string, channels ...string) (bot *MoonBot, err error) {
	conn, err := irc.Dial(server, Nick, Real, false)
	if err != nil {
		return
	}
	conn.SendRaw("PASS " + password + "\n")

	bot = &MoonBot{
		Conn: conn,
	}

	go func(b *MoonBot) {
		for msg := range bot.Conn.Read {
			if msg.Trailing == ":quit" {
				b.Quit()
			}
		}
	}(bot)

	return
}

func (b *MoonBot) Say(channel, message string) {
	b.Conn.SendMessage(channel, message)
}

func (b *MoonBot) Quit() {
	l.Println("Got quit signal.")
	b.Say("#moonbot", "Quitting.")
	quitSig <- 0
}
