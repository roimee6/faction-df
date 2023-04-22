package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

func ExistFaction(name string) bool {
	_, ok := util.Factions[strings.ToLower(name)]
	if ok {
		return true
	} else {
		return false
	}
}

func GetFactionMembers(key string) []string {
	var arr []string

	faction := util.Factions[key]

	list := faction.Members
	leader := list.Leader

	arr = append(arr, leader)
	members := append(list.Officers, list.Members...)

	for _, p := range members {
		arr = append(arr, p)
	}
	return arr
}

func HasFaction(player *player.Player) bool {
	_ = GetFactionRankOnline(player)

	user, err := session.GetUser(player)
	if err != nil {
		return false
	}

	if user.Data.Faction == nil {
		return false
	} else {
		return true
	}
}

func GetFactionRankOnline(p *player.Player) *string {
	user, err := session.GetUser(p)
	if err != nil {
		return nil
	}

	if user.Data.Faction == nil {
		return nil
	} else if !ExistFaction(*user.Data.Faction) {
		user.Data.Faction = nil
		return nil
	}

	rank := GetFactionRank(*user.Data.Faction, p.Name())

	if rank == nil {
		user.Data.Faction = nil
	}
	return rank
}

func GetOnlineFactionMembers(faction string) []*player.Player {
	var arr []*player.Player

	members := GetFactionMembers(strings.ToLower(faction))

	for _, name := range members {
		if p, ok := util.Server.PlayerByName(name); ok {
			arr = append(arr, p)
		}
	}
	return arr
}

func BroadcastFactionMessage(faction string, message string) {
	for _, p := range GetOnlineFactionMembers(faction) {
		p.Message(message)
	}
}

func GetFactionRank(faction string, player string) *string {
	members := util.Factions[strings.ToLower(faction)].Members

	if members.Leader == player {
		rank := "leader"
		return &rank
	} else if util.InArray(player, members.Officers) {
		rank := "officier"
		return &rank
	} else if util.InArray(player, members.Members) {
		rank := "member"
		return &rank
	} else {
		return nil
	}
}
