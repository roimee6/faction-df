package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/roimee6/Faction/server/handler"
	"strings"
)

type Say struct {
	Message cmd.Varargs `cmd:"message"`
}

func (s Say) Run(_ cmd.Source, _ *cmd.Output) {
	msg := strings.TrimSpace(string(s.Message))

	if msg == "" {
		return
	}

	_, err := chat.Global.WriteString(msg)

	if err != nil {
		return
	}
}

func (Say) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}
