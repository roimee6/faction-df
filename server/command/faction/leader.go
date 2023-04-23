package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type LeaderByTarget struct {
	Sub     cmd.SubCommand `cmd:"leader"`
	Targets []cmd.Target   `cmd:"joueur"`
}

type LeaderByString struct {
	Sub    cmd.SubCommand `cmd:"leader"`
	Target string         `cmd:"joueur"`
}

func (p LeaderByString) Run(source cmd.Source, _ *cmd.Output) {
	name := p.Target
	Leader(name, source)
}

func (p LeaderByTarget) Run(source cmd.Source, _ *cmd.Output) {
	targets := p.Targets

	if len(targets) < 1 {
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		return
	}

	name := target.Name()
	Leader(name, source)
}

func Leader(name string, source cmd.Source) {
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
	} else if *rank != "leader" {
		sender.Message("Vous ne pouvez pas supprimer la faction si vous n'êtes pas le chef")
		return
	} else if !util.InArray(name, handler.GetFactionMembers(*faction)) || targetRank == nil {
		sender.Message("Ce joueur n'est pas dans votre faction, reessayez en vérifiant les majuscules !")
		return
	} else if name == sender.Name() {
		sender.Message("Vous ne pouvez pas vous promouvoir vous-même !")
		return
	}

	switch *targetRank {
	case "officer":
		handler.Factions[*faction].Members.Officers = util.RemoveElementFromArray(handler.Factions[*faction].Members.Officers, name)
		break
	case "member":
		handler.Factions[*faction].Members.Members = util.RemoveElementFromArray(handler.Factions[*faction].Members.Members, name)
		break
	}

	handler.Factions[*faction].Members.Leader = name
	handler.Factions[*faction].Members.Officers = append(handler.Factions[*faction].Members.Officers, sender.Name())

	sender.Messagef("Vous venez de promouvoir %s chef !", name)
}
