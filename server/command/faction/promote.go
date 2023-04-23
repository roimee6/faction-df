package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type PromoteByTarget struct {
	Sub     cmd.SubCommand `cmd:"promote"`
	Targets []cmd.Target   `cmd:"joueur"`
}

type PromoteByString struct {
	Sub    cmd.SubCommand `cmd:"promote"`
	Target string         `cmd:"joueur"`
}

func (p PromoteByString) Run(source cmd.Source, _ *cmd.Output) {
	name := p.Target
	promote(name, source)
}

func (p PromoteByTarget) Run(source cmd.Source, _ *cmd.Output) {
	targets := p.Targets

	if len(targets) < 1 {
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		return
	}

	name := target.Name()
	promote(name, source)
}

func promote(name string, source cmd.Source) {
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
	} else if *targetRank == "officer" || *targetRank == "leader" {
		sender.Message("Ce joueur ne peut pas être promu car il est déjà officer")
		return
	}

	handler.Factions[*faction].Members.Members = util.RemoveElementFromArray(handler.Factions[*faction].Members.Members, name)
	handler.Factions[*faction].Members.Officers = append(handler.Factions[*faction].Members.Officers, name)

	sender.Messagef("Vous venez de promouvoir %s en officer !", name)
}
