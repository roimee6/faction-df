package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"strconv"
)

type Tl struct{}

func (Tl) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	x := strconv.Itoa(int(sender.Position().X()))
	y := strconv.Itoa(int(sender.Position().Y()))
	z := strconv.Itoa(int(sender.Position().Z()))

	handler.BroadcastFactionMessage(*user.Data.Faction, "[F] "+sender.Name()+" » X: "+x+", Y: "+y+", Z: "+z)
}
