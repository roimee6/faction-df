package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
	"time"
)

type Weather struct {
	Weather  weatherType        `cmd:"weather"`
	Duration cmd.Optional[uint] `cmd:"duration"`
}

func (c Weather) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	w := sender.World()
	t := time.Duration(c.Duration.LoadOr(60)) * time.Second

	switch strings.ToLower(string(c.Weather)) {
	case "clear":
		w.StopRaining()
		w.StopThundering()
		sender.Message("Changing to clear weather")
	case "rain":
		w.StartRaining(t)
		sender.Message("Changing to rainy weather")
	case "thunder":
		w.StartRaining(t)
		w.StartThundering(t)
		sender.Message("Changing to rain and thunder")
	}
}

type weatherType string

func (w weatherType) Type() string {
	return "weather"
}

func (w weatherType) Options(_ cmd.Source) []string {
	return []string{
		"clear", "rain", "thunder",
	}
}
