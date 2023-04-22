package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type SetTimeInt struct {
	Sub    cmd.SubCommand `cmd:"set"`
	Amount int            `name:"amount"`
}

type Add struct {
	Sub    cmd.SubCommand `cmd:"add"`
	Amount int            `name:"amount"`
}

func (t Add) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	w := sender.World()
	w.SetTime(w.Time() + t.Amount)

	sender.Messagef("Added %s to the time", t.Amount)
}

type SetTimeString struct {
	Sub  cmd.SubCommand `cmd:"set"`
	Time spec           `name:"time"`
}

func setTime(source cmd.Source, _ *cmd.Output, t int) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	w := sender.World()
	w.SetTime(t)

	sender.Messagef("Set the time to %s", t)
}

func (t SetTimeInt) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, t.Amount)
}

func (t SetTimeString) Run(source cmd.Source, output *cmd.Output) {
	tf := map[spec]int64{
		"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
	}[t.Time]

	setTime(source, output, int(tf))
}

type spec string

func (spec) Type() string {
	return "TimeSpec"
}

func (spec) Options(_ cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}
