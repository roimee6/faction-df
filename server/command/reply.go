package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
	"strings"
)

type Reply struct {
	Message cmd.Varargs `cmd:"message"`
}

func (r Reply) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if user.Data.LastReply == nil {
		sender.Message("Vous n'avez personne à qui répondre !")
		return
	}

	username := user.Data.LastReply
	message := strings.TrimSpace(string(r.Message))

	sender.ExecuteCommand("/msg " + *username + " " + message)
}
