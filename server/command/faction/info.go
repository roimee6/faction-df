package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
)

type Info struct {
	Sub cmd.SubCommand `cmd:"info"`
}

func (i Info) Run(source cmd.Source, _ *cmd.Output) {
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

	faction := user.Data.Faction

	sender.Message("Faction: " + *faction)
	sender.Messagef("%d", handler.Factions[*faction])
}
