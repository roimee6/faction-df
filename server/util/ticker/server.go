package ticker

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/util"
	"time"
)

var (
	lastPosition map[string]mgl64.Vec3
	Combat       []string
)

func InitBaseTask() {
	lastPosition = make(map[string]mgl64.Vec3)

	go func() {
		for {
			server := util.Server
			players := server.Players()

			for _, p := range players {
				updatePlayer(p)
			}

			time.Sleep(time.Second)
		}
	}()
}

func updatePlayer(p *player.Player) {
	if util.InArray(p.Name(), Combat) {
		if handler.InCooldown(p, "combat") {
			if handler.InsideZone(p.Position(), "spawn") {
				pos, ok := lastPosition[p.Name()]

				if ok {
					p.Teleport(pos)
				}
			}
		} else {
			Combat = util.RemoveElementFromArray(Combat, p.Name())
			p.Message("§cVous n'êtes plus en combat !")
		}
	}

	if !handler.InsideZone(p.Position(), "spawn") {
		lastPosition[p.Name()] = p.Position()
	}

	handler.UpdateScoreboard(p)
}
