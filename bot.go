package main

import (
	"github.com/museun/irc"
	"log"
	"strings"
)

type MoonBot struct {
	Conn *irc.Conn

	l *log.Logger
}

const (
	Nick = "MoonBot"
	Real = "Duo's Bot"
)

func NewBot(server, password string, logger *log.Logger, channels ...string) (bot *MoonBot, err error) {
	conn, err := irc.Dial(server, Nick, Real, false)
	if err != nil {
		return
	}
	conn.SendRaw("PASS " + password + "\n")

	bot = &MoonBot{
		Conn: conn,
		l:    logger,
	}

	go func(b *MoonBot) {
		for msg := range bot.Conn.Read {
			command := strings.SplitN(msg.Trailing, " ", 3)
			if len(command) >= 2 && command[0] == Nick+":" {
				var args string
				channel := "#moonbot"
				if len(command) >= 3 {
					args = command[2]
				}
				switch strings.ToLower(command[1]) {
				case "version":
					b.Say("#moonbot", "Version "+Version+minversion)

				case "quit":
					b.l.Println("Quitting")
					b.Quit()

				case "echo":
					b.l.Println("Echoing", args)
					b.Say(channel, args)
				}
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
