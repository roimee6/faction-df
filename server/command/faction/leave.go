package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type Leave struct {
	Sub cmd.SubCommand `cmd:"leave"`
}

func (l Leave) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	rank := handler.GetFactionRankOnline(sender)

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if *rank == "leader" {
		sender.Message("Vous ne pouvez pas quitter votre faction car vous êtes le leader !")
		return
	}

	faction := user.Data.Faction

	switch *rank {
	case "officer":
		handler.Factions[*faction].Members.Officers = util.RemoveElementFromArray(handler.Factions[*faction].Members.Officers, sender.Name())
		break
	case "member":
		handler.Factions[*faction].Members.Members = util.RemoveElementFromArray(handler.Factions[*faction].Members.Members, sender.Name())
		break
	}

	user.Data.Faction = nil
	sender.Message("Vous venez de quitter votre faction !")

	handler.UpdateNameTag(sender)
}
