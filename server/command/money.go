package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
)

type Money struct {
	Player cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (m Money) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	targets := m.Player.LoadOr([]cmd.Target{})

	if len(targets) < 1 {
		user, err := session.GetUser(sender)
		if err != nil {
			return
		}

		sender.Messagef("Vous possèdez %s pièces !", user.Data.Money)
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		sender.Message("Le joueur spécifié n'est pas connecté !")
		return
	}

	user, err := session.GetUser(target)
	if err != nil {
		return
	}

	sender.Messagef("Le joueur %s possède %s pièces !", target.Name(), user.Data.Money)
}
