package command

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/handler"
)

type TeleportXYZ struct {
	X float64 `cmd:"x"`
	Y float64 `cmd:"y"`
	Z float64 `cmd:"z"`
}

func (t TeleportXYZ) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)

	if !ok {
		return
	}

	sender.Teleport(mgl64.Vec3{t.X, t.Y, t.Z})
	sender.Messagef("Teleported to X: %s Y: %s Z: %s", int(t.X), int(t.Y), int(t.Z))
}

type TeleportPlayer struct {
	Player  []cmd.Target               `cmd:"player"`
	Targets cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (t TeleportPlayer) Run(source cmd.Source, output *cmd.Output) {
	if len(t.Player) < 1 {
		output.Error("Usage: /teleport <Player: target> [Target: target]")
	}

	targets := t.Targets.LoadOr(nil)

	for _, target := range t.Player {
		if tp1, ok := target.(*player.Player); ok {
			if len(targets) < 1 {
				if p, ok := source.(*player.Player); ok {
					t.TeleportAnotherPlayer(p, tp1)
				}
			} else if tp2, ok := targets[0].(*player.Player); ok {
				t.TeleportPlayerToAnotherPlayer(tp1, tp2)
				output.Printf("%s has been teleported to %s.", tp1.Name(), tp2.Name())
			} else {
				output.Errorf("Second target is not a player!")
			}
		} else {
			output.Errorf("First target is not a player!")
		}
	}
}

func (TeleportPlayer) TeleportAnotherPlayer(p, to *player.Player) {
	p.Teleport(to.Position())
	p.Message(fmt.Sprintf("You teleported to %s.", to.Name()))
}

func (TeleportPlayer) TeleportPlayerToAnotherPlayer(to, p *player.Player) {
	to.Teleport(p.Position())

	p.Message(fmt.Sprintf("%s has been teleported to your side.", to.Name()))
	to.Message(fmt.Sprintf("You have been teleported here by %s.", p.Name()))
}

func (TeleportPlayer) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}

func (TeleportXYZ) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}
