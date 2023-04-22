package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

type Create struct {
	Sub  cmd.SubCommand `cmd:"create"`
	Name string         `cmd:"name"`
}

func (c Create) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	name := c.Name

	if handler.HasFaction(sender) {
		sender.Message("Vous avez déjà une faction !")
		return
	} else if handler.ExistFaction(name) {
		sender.Message("Cette faction existe déjà !")
		return
	} else if len(name) > 16 || !util.IsStringAlphanumeric(name) {
		sender.Message("Le nom de la faction est invalide")
		return
	}

	util.Factions[strings.ToLower(name)] = &util.Faction{
		Name: name,
		Members: util.FactionMembers{
			Leader: sender.Name(),
		},
	}

	user.Data.Faction = &name
	sender.Message("Vous venez de créer la faction " + name)
}
