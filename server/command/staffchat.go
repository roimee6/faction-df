package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
)

type StaffChat struct{}

func (s StaffChat) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if user.Data.Coordinates {
		sender.Message("Vous venez de désactiver le staffchat !")
	} else {
		sender.Message("Vous venez d'activer le staffchat !")
	}

	user.Data.StaffChat = !user.Data.Coordinates
}

func (StaffChat) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "guide")
}
