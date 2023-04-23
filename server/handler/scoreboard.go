package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strconv"
	"strings"
)

func UpdateScoreboard(player *player.Player) {
	user, err := session.GetUser(player)

	if err != nil || !user.Data.Scoreboard {
		return
	}

	s := scoreboard.New("Nitro")
	s.Set(0, "§f")
	s.Set(5, "§r")
	s.Set(6, "§l§9Serveur")
	s.Set(7, "§fConnectés: §9"+strconv.Itoa(len(util.Server.Players())))
	s.Set(8, "§f§fVoteParty: §90")
	s.Set(9, "§l")
	s.Set(10, "     §7nitrofaction.fr    ")

	rank := "Joueur"
	faction := "Aucune"

	if player.Name() == user.Data.DisplayName {
		rank = user.Data.Rank
		rank = strings.ToUpper(rank[:1]) + rank[1:]
	}

	if HasFaction(player) {
		faction = Factions[*user.Data.Faction].Name
	}

	s.Set(1, "§l§9"+user.Data.DisplayName)
	s.Set(2, "§fGrade: §9"+rank)
	s.Set(3, "§fFaction: §9"+faction)
	s.Set(4, "§fPieces: §9"+util.FormatInt(user.Data.Money))

	s.RemovePadding()
	player.SendScoreboard(s)
}
