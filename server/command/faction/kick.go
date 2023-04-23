package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type KickByTarget struct {
	Sub     cmd.SubCommand `cmd:"kick"`
	Targets []cmd.Target   `cmd:"joueur"`
}

type KickByString struct {
	Sub    cmd.SubCommand `cmd:"kick"`
	Target string         `cmd:"joueur"`
}

func (p KickByString) Run(source cmd.Source, _ *cmd.Output) {
	name := p.Target
	kick(name, source)
}

func (p KickByTarget) Run(source cmd.Source, _ *cmd.Output) {
	targets := p.Targets

	if len(targets) < 1 {
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		return
	}

	name := target.Name()
	kick(name, source)
}

func kick(name string, source cmd.Source) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	faction := user.Data.Faction

	rank := handler.GetFactionRankOnline(sender)
	targetRank := handler.GetFactionRank(*faction, name)

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if *rank == "member" {
		sender.Message("Vous ne pouvez pas supprimer la faction si vous n'êtes pas le chef")
		return
	} else if !util.InArray(name, handler.GetFactionMembers(*faction)) || targetRank == nil {
		sender.Message("Ce joueur n'est pas dans votre faction, reessayez en vérifiant les majuscules !")
		return
	} else if name == sender.Name() {
		sender.Message("Vous ne pouvez pas vous kick vous-même !")
		return
	} else if rank == targetRank {
		sender.Message("Vous ne pouvez pas kick un membre du même rang que vous !")
		return
	} else if *targetRank == "leader" {
		sender.Message("Vous ne pouvez pas kick le chef de la faction !")
		return
	}

	if target, ok := util.Server.PlayerByName(name); ok {
		targetUser, err := session.GetUser(target)

		if err == nil {
			targetUser.Data.Faction = nil
			handler.UpdateNameTag(target)
		}
	}

	switch *targetRank {
	case "officer":
		handler.Factions[*faction].Members.Officers = util.RemoveElementFromArray(handler.Factions[*faction].Members.Officers, name)
		break
	case "member":
		handler.Factions[*faction].Members.Members = util.RemoveElementFromArray(handler.Factions[*faction].Members.Members, name)
		break
	}

	sender.Message("Vous avez kick %s de votre faction !", name)
}
