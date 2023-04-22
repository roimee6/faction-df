package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
)

type Invite struct {
	Sub    cmd.SubCommand `cmd:"invite"`
	Player []cmd.Target   `cmd:"player"`
}

func (i Invite) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	if len(i.Player) < 1 {
		sender.Message("Vous devez spécifier un joueur !")
		return
	}

	target, ok := i.Player[0].(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	rank := handler.GetFactionRankOnline(sender)
	faction := user.Data.Faction

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if *rank == "member" {
		sender.Message("Vous n'avez pas la permission d'inviter des joueurs !")
		return
	} else if handler.HasFaction(target) {
		sender.Message("Ce joueur a déjà une faction !")
		return
	} else if len(handler.GetFactionMembers(*faction)) >= 20 {
		sender.Message("Votre faction est pleine !")
		return
	}

	targetUser, err := session.GetUser(target)
	if err != nil {
		return
	}

	targetUser.Data.Invites = append(targetUser.Data.Invites, *faction)

	sender.Message("Vous venez d'inviter " + target.Name() + " dans votre faction !")
	target.Message("Vous avez été invité dans la faction " + *faction + " !")
}
