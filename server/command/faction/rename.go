package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

type Rename struct {
	Sub  cmd.SubCommand `cmd:"rename"`
	Name string         `cmd:"name"`
}

func (r Rename) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	name := r.Name
	rank := handler.GetFactionRankOnline(sender)

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if *rank != "leader" {
		sender.Message("Vous ne pouvez pas supprimer la faction si vous n'êtes pas le chef")
		return
	} else if handler.ExistFaction(name) {
		sender.Message("Cette faction existe déjà !")
		return
	} else if len(name) > 16 || !util.IsStringAlphanumeric(name) {
		sender.Message("Le nom de la faction est invalide")
		return
	} else if handler.InCooldown(sender, "rename") {
		sender.Messagef("Vous devez attendre %s avant de pouvoir renommer votre faction", util.FormatSeconds(handler.GetCooldownTime(sender, "rename"), 0))
		return
	}

	faction := *user.Data.Faction

	members := handler.GetFactionMembers(faction)
	newName := strings.ToLower(name)

	oldFac := handler.Factions[faction]
	oldFac.Name = name

	handler.Factions[newName] = oldFac
	delete(handler.Factions, faction)

	for _, member := range members {
		if p, ok := util.Server.PlayerByName(member); ok {
			user, err := session.GetUser(p)
			if err != nil {
				continue
			}

			user.Data.Faction = &newName
			handler.UpdateNameTag(p)
		}

		xuid := handler.GetOfflinePlayerXuid(member)

		data := session.ParseData(xuid)
		data.Faction = &newName

		session.SaveData(xuid, data)
	}

	handler.SetCooldown(sender, "rename", 5*60, []string{})
}
