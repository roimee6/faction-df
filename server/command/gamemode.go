package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"strings"
)

type GameMode interface {
	Run(cmd.Source, *cmd.Output)
	GetGameMode() interface{}
	GetTargets() []cmd.Target
}

type GameModeString struct {
	GameMode gameMode                   `cmd:"gameMode"`
	Targets  cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (g GameModeString) Run(source cmd.Source, _ *cmd.Output) {
	updateGameMode(source, &g)
}

func (g GameModeString) GetGameMode() interface{} {
	return strings.ToLower(string(g.GameMode))
}

func (g GameModeString) GetTargets() []cmd.Target {
	return g.Targets.LoadOr(nil)
}

type GameModeInt struct {
	GameMode int                        `cmd:"gamemode"`
	Targets  cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (g GameModeInt) Run(source cmd.Source, _ *cmd.Output) {
	updateGameMode(source, &g)
}

func (g GameModeInt) GetGameMode() interface{} {
	return g.GameMode
}

func (g GameModeInt) GetTargets() []cmd.Target {
	return g.Targets.LoadOr(nil)
}

func updateGameMode(source cmd.Source, g GameMode) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	var mode world.GameMode
	gameMode := g.GetGameMode()

	switch gameMode {
	case "survival", 0, "s":
		_, mode = "survival", world.GameModeSurvival
	case "creative", 1, "c":
		_, mode = "creative", world.GameModeCreative
	case "adventure", 2, "a":
		_, mode = "adventure", world.GameModeAdventure
	case "spectator", 3, "sp":
		_, mode = "spectator", world.GameModeSpectator
	}

	targets := g.GetTargets()

	if len(targets) > 1 {
		sender.Message("Hey")
		return
	}

	if len(targets) == 1 {
		target, ok := targets[0].(*player.Player)

		if !ok {
			sender.Message("Le joueur mentionné n'est pas connecté au serveur")
			return
		}

		target.SetGameMode(mode)
		sender.Message("update other")
		return
	}

	sender.SetGameMode(mode)
	sender.Message("update")
}

type gameMode string

func (gameMode) Type() string {
	return "GameMode"
}

func (gameMode) Options(cmd.Source) []string {
	return []string{
		"survival", "s",
		"creative", "c",
		"adventure", "a",
		"spectator", "sp",
	}
}
