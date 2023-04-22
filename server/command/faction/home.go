package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strconv"
	"strings"
	"time"
)

type Home struct {
	Sub cmd.SubCommand `cmd:"home"`
}

func (h Home) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	} else if session.InCooldown(sender, "combat") {
		sender.Message("Vous ne pouvez pas faire ça en combat !")
		return
	} else if session.InCooldown(sender, "teleportation") {
		sender.Message("Vous ne pouvez pas faire ça si vous êtes en téléportation !")
		return
	}

	faction := user.Data.Faction
	fac := util.Factions[*faction]

	if fac.Home == nil {
		sender.Message("Votre faction n'a pas de base !")
		return
	}

	home := strings.Split(*fac.Home, ":")

	x, _ := strconv.Atoi(home[0])
	y, _ := strconv.Atoi(home[1])
	z, _ := strconv.Atoi(home[2])

	pos := mgl64.Vec3{
		float64(x), float64(y), float64(z),
	}

	t := handler.GetTpTime(sender)
	if t > 0 {
		sender.AddEffect(effect.New(effect.Blindness{}, 1, time.Duration(t+1)*time.Second).WithoutParticles())
	}

	session.SetCooldown(sender, "teleportation", int64(t), []string{handler.GetPlace(sender)})
	handler.CreateTeleportatationTicker(sender, pos)

	sender.Message("Vous allez être téléporté à la base !")
}
