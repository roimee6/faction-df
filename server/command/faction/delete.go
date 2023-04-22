package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type Delete struct {
	Sub cmd.SubCommand `cmd:"delete"`
}

func (d Delete) Run(source cmd.Source, _ *cmd.Output) {
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
	} else if *rank != "leader" {
		sender.Message("Vous ne pouvez pas supprimer la faction si vous n'êtes pas le chef")
		return
	}

	handler.BroadcastFactionMessage(*user.Data.Faction, "La faction dont vous êtiez n'existe désormais plus")

	for _, member := range handler.GetOnlineFactionMembers(*user.Data.Faction) {
		memberUser, err := session.GetUser(member)
		if err != nil {
			continue
		}

		memberUser.Data.Faction = nil
	}

	delete(util.Factions, *user.Data.Faction)
}
