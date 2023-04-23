package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"strings"
)

type Sudo struct {
	Targets []cmd.Target `cmd:"target"`
	Message cmd.Varargs  `cmd:"message"`
}

func (s Sudo) Run(source cmd.Source, _ *cmd.Output) {
	sender := source.(*player.Player)

	var names []string
	var command bool

	if len(s.Targets) < 1 {
		sender.Message("No targets specified!")
		return
	}

	msg := strings.TrimSpace(string(s.Message))

	if strings.HasPrefix(msg, "/") {
		command = true
	}

	target := s.Targets[0]

	if tar, ok := target.(*player.Player); ok {
		if command {
			tar.ExecuteCommand(msg)
		} else {
			tar.Chat(msg)
		}

		names = append(names, tar.Name())
	}

	sender.Messagef("Message sent as %s !", strings.Join(names, ", "))
}

func (Sudo) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}
