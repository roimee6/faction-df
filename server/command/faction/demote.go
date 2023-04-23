package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type DemoteByTarget struct {
	Sub     cmd.SubCommand `cmd:"demote"`
	Targets []cmd.Target   `cmd:"player"`
}

type DemoteByString struct {
	Sub    cmd.SubCommand `cmd:"demote"`
	Target string         `cmd:"player"`
}

func (p DemoteByString) Run(source cmd.Source, _ *cmd.Output) {
	name := p.Target
	demote(name, source)
}

func (p DemoteByTarget) Run(source cmd.Source, _ *cmd.Output) {
	targets := p.Targets

	if len(targets) < 1 {
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		return
	}

	name := target.Name()
	demote(name, source)
}

func demote(name string, source cmd.Source) {
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
	} else if *targetRank == "member" {
		sender.Message("Ce joueur est déjà un membre !")
		return
	}

	handler.Factions[*faction].Members.Officers = util.RemoveElementFromArray(handler.Factions[*faction].Members.Officers, name)
	handler.Factions[*faction].Members.Members = append(handler.Factions[*faction].Members.Members, name)

	sender.Messagef("Vous avez rétrogradé %s à membre !", name)
}
