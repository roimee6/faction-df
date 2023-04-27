package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/session"
	"strconv"
)

type Pay struct {
	Targets []cmd.Target `cmd:"target"`
	Amount  int          `cmd:"amount"`
}

func (p Pay) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	if len(p.Targets) < 1 {
		sender.Message("Vous devez spécifier un joueur !")
		return
	}

	target, ok := p.Targets[0].(*player.Player)

	if !ok {
		sender.Message("Le joueur indiqué n'est pas connecté sur le serveur")
		return
	}

	senderUser, err := session.GetUser(sender)
	if err != nil {
		return
	}

	targetUser, err := session.GetUser(target)
	if err != nil {
		return
	}

	if p.Amount < 1 {
		sender.Message("Vous devez spécifier un montant !")
		return
	} else if senderUser.Data.Money < p.Amount {
		sender.Message("Vous n'avez pas assez d'argent !")
		return
	}

	senderUser.Data.Money -= p.Amount
	targetUser.Data.Money += p.Amount

	sender.Message("Vous venez de donner " + strconv.Itoa(p.Amount) + " pièces à " + target.Name())
	target.Message("Vous venez de recevoir " + strconv.Itoa(p.Amount) + " pièces de la part de " + sender.Name())
}
