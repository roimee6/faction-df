package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/util"
	"github.com/roimee6/Faction/server/util/ticker"
	"time"
)

type Spawn struct{}

func (Spawn) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	if handler.InCooldown(sender, "combat") {
		sender.Message("Vous ne pouvez pas faire ça en combat !")
		return
	} else if handler.InCooldown(sender, "teleportation") {
		sender.Message("Vous ne pouvez pas faire ça si vous êtes en téléportation !")
		return
	}

	t := handler.GetTpTime(sender)
	if t > 0 {
		sender.AddEffect(effect.New(effect.Blindness{}, 1, time.Duration(t+1)*time.Second).WithoutParticles())
	}

	handler.SetCooldown(sender, "teleportation", int64(t), []string{handler.GetPlace(sender)})
	ticker.Teleportation(sender, util.Server.World().Spawn().Vec3())

	sender.Message("Vous allez être téléporté à la base !")
}
