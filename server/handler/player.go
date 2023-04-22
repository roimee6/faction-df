package handler

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/session"
	"strconv"
	"time"
)

func CreateTeleportatationTicker(player *player.Player, pos mgl64.Vec3) {
	go func() {
		for {
			data := session.GetCooldownData(player, "teleportation")

			t, _ := strconv.Atoi(data[0])
			t = t - int(time.Now().Unix())

			if data[1] != GetPlace(player) {
				player.Message("Vous avez bougé, la téléportation a été annulée !")
				cancel(player)
				break
			} else if !session.InCooldown(player, "teleportation") {
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
	session.RemoveCooldown(player, "teleportation")
}

func GetTpTime(player *player.Player) int {
	if player.GameMode() == world.GameModeCreative || player.GameMode() == world.GameModeAdventure {
		return -1
	} else {
		return 5
	}
}

func GetPlace(player *player.Player) string {
	pos := player.Position()
	return strconv.Itoa(int(pos.X())) + ":" + strconv.Itoa(int(pos.Y())) + ":" + strconv.Itoa(int(pos.Z()))
}
