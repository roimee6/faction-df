package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
)

type Xyz struct{}

func (c Xyz) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	user, err := session.GetUser(sender)

	if err != nil {
		return
	}

	if user.Data.Coordinates {
		sender.HideCoordinates()
		sender.Message("Vous venez de désactiver les coordonnées !")
	} else {
		sender.ShowCoordinates()
		sender.Message("Vous venez d'activer les coordonnées !")
	}

	user.Data.Coordinates = !user.Data.Coordinates
}
