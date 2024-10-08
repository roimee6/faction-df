package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
)

type Delhome struct {
	Sub cmd.SubCommand `cmd:"delhome"`
}

func (s Delhome) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	rank := handler.GetFactionRankOnline(sender)
	faction := user.Data.Faction

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if *rank == "member" {
		sender.Message("Vous n'avez pas la permission d'inviter des joueurs !")
		return
	}

	handler.Factions[*faction].Home = nil
	sender.Message("Vous venez de supprimer votre base !")
}
