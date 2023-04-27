package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strconv"
)

type Ping struct {
	Targets cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (p Ping) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	targets := p.Targets.LoadOr([]cmd.Target{})

	if len(targets) < 1 {
		latency := sender.Latency().Milliseconds() * 2
		sender.Message("Vous possèdez " + strconv.Itoa(int(latency)) + " de ping")
		return
	}

	target, ok := targets[0].(*player.Player)

	if !ok {
		sender.Message("Le joueur indiqué n'est pas connecté sur le serveur")
		return
	}

	latency := sender.Latency().Milliseconds() * 2
	sender.Message("Le joueur " + target.Name() + " possède " + strconv.Itoa(int(latency)) + " de ping")
}
