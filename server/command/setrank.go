package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
)

type SetRankByString struct {
	Target string `cmd:"joueur"`
	Rank   rank   `cmd:"rank"`
}

type SetRankByTarget struct {
	Targets []cmd.Target `cmd:"joueur"`
	Rank    rank         `cmd:"rank"`
}

func (s SetRankByString) Run(_ cmd.Source, output *cmd.Output) {
	name := s.Target
	setRank(name, string(s.Rank), output)
}

func (s SetRankByTarget) Run(_ cmd.Source, output *cmd.Output) {
	targets := s.Targets

	if len(targets) < 1 {
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		return
	}

	name := target.Name()
	setRank(name, string(s.Rank), output)
}

func setRank(name string, rank string, output *cmd.Output) {
	xuid := handler.GetOfflinePlayerXuid(name)

	if xuid == "" {
		output.Errorf("Le joueur %s n'existe pas !", name)
		return
	}

	if p, ok := util.Server.PlayerByName(name); ok {
		user, err := session.GetUser(p)
		if err != nil {
			return
		}

		user.Data.Rank = rank
		handler.UpdateNameTag(p)
	} else {
		data := session.ParseData(xuid)
		data.Rank = rank
		session.SaveData(xuid, data)
	}

	output.Printf("Le joueur %s a recu le grade %s !", name, rank)
}

type rank string

func (rank) Type() string {
	return "rank"
}

func (rank) Options(cmd.Source) []string {
	return util.RanksList
}

func (SetRankByTarget) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}

func (SetRankByString) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}
