package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
)

type Sudo struct {
	Targets []cmd.Target `cmd:"target"`
	Message cmd.Varargs  `cmd:"message"`
}

func (s Sudo) Run(source cmd.Source, _ *cmd.Output) {
	p := source.(*player.Player)

	var names []string
	var command bool

	if len(s.Targets) < 1 {
		p.Message("No targets specified!")
		return
	}

	msg := strings.TrimSpace(string(s.Message))

	if strings.HasPrefix(msg, "/") {
		command = true
	}

	for _, target := range s.Targets {
		if tar, ok := target.(*player.Player); ok {
			if command {
				tar.ExecuteCommand(msg)
			} else {
				tar.Chat(msg)
			}

			names = append(names, tar.Name())
		}
	}

	p.Messagef("Message sent as %s !", strings.Join(names, ", "))
}
