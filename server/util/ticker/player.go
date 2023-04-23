package ticker

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/handler"
	"strconv"
	"time"
)

func Teleportation(player *player.Player, pos mgl64.Vec3) {
	go func() {
		for {
			data := handler.GetCooldownData(player, "teleportation")
			t := handler.GetCooldownTime(player, "teleportation")

			if data[1] != handler.GetPlace(player) {
				player.Message("Vous avez bougé, la téléportation a été annulée !")
				cancel(player)
				break
			} else if !handler.InCooldown(player, "teleportation") {
				player.Teleport(pos)
				cancel(player)
				break
			}

			player.SendTip("Téléportation dans: " + strconv.Itoa(t))
			player.PlaySound(sound.Click{})

			time.Sleep(time.Second)
		}
	}()
}

func cancel(player *player.Player) {
	player.RemoveEffect(effect.Blindness{})
	handler.RemoveCooldown(player, "teleportation")
}
