package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

type Accept struct {
	Sub     cmd.SubCommand       `cmd:"accept"`
	Faction cmd.Optional[string] `cmd:"faction"`
}

func (a Accept) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	faction := a.Faction.LoadOr("")

	if handler.HasFaction(sender) {
		sender.Message("Vous avez déjà une faction !")
		return
	} else if len(user.Data.Invites) == 0 {
		sender.Message("Vous n'avez aucune invitation !")
		return
	} else if faction == "" && len(user.Data.Invites) != 1 {
		sender.Message("Vous possèdez plus qu'une invitation de faction, merci d'écrire la faction que vous voulez rejoindre")
		return
	}

	if faction == "" {
		faction = user.Data.Invites[0]
	}

	faction = strings.ToLower(faction)

	if !handler.ExistFaction(faction) {
		sender.Message("La faction que vous vouliez rejoindre n'existe plus")
		return
	} else if !util.InArray(faction, user.Data.Invites) {
		sender.Message("Vous n'avez pas été invité dans cette faction")
		return
	} else if len(handler.GetFactionMembers(faction)) >= 20 {
		sender.Message("Cette faction est pleine !")
		return
	}

	user.Data.Invites = []string{}
	user.Data.Faction = &faction

	util.Factions[faction].Members.Members = append(util.Factions[faction].Members.Members, sender.Name())
	sender.Messagef("Vous avez rejoint la faction %s !", faction)
}
