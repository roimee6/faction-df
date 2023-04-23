package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
)

type Chat struct {
	Sub cmd.SubCommand `cmd:"chat"`
}

func (c Chat) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	}

	if user.Data.FactionChat {
		sender.Message("Vous venez d'enlever le chat de faction !")
	} else {
		sender.Message("Vous venez de mettre le chat de faction !")
	}

	user.Data.FactionChat = !user.Data.FactionChat
}
