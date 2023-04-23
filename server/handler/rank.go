package handler

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

func UpdateNameTag(p *player.Player) {
	name := p.Name()
	rank := "joueur"

	user, err := session.GetUser(p)
	if err != nil {
		return
	}

	if name == user.Data.DisplayName {
		rank = user.Data.Rank
	}

	replace := GetRankValue(rank, "gamertag")
	nametag := SetReplace(replace, p, "")

	p.SetNameTag(nametag)
}

func GetMessage(p *player.Player, message string) string {
	name := p.Name()
	rank := "joueur"

	user, err := session.GetUser(p)
	if err != nil {
		return ""
	}

	if name == user.Data.DisplayName {
		rank = user.Data.Rank
	}

	prefix := GetRankValue(rank, "chat")
	replace := SetReplace(prefix, p, message)

	return replace
}

func GetRankValue(rank string, key string) string {
	rank = strings.ToLower(rank)
	return util.Ranks[rank][key]
}

func SetReplace(replace string, p *player.Player, msg string) string {
	fac := HasFaction(p)

	user, err := session.GetUser(p)
	if err != nil {
		return ""
	}

	var faction string

	if fac {
		faction = Factions[*user.Data.Faction].Name
	} else {
		faction = "..."
	}

	return strings.NewReplacer(
		"{name}", user.Data.DisplayName,
		"{fac}", faction,
		"{msg}", msg,
	).Replace(replace)
}

func hasRank(rank string, checkRank string) bool {
	return util.IndexOf(util.RanksList, rank) >= util.IndexOf(util.RanksList, checkRank)
}

func HasRankPermission(source cmd.Source, rank string) bool {
	sender, ok := source.(*player.Player)

	if !ok {
		return true
	} else {
		user, err := session.GetUser(sender)
		if err != nil {
			return false
		}

		return hasRank(user.Data.Rank, rank)
	}
}
