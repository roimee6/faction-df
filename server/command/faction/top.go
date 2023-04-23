package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Top struct {
	Sub cmd.SubCommand `cmd:"top"`
}

func (t Top) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	sender.ExecuteCommand("/top faction")
}
