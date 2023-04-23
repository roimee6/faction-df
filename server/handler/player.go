package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"math"
	"strconv"
	"strings"
	"time"
)

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

func GetOfflinePlayerXuid(name string) string {
	xuid, ok := Players["xuid"][strings.ToLower(name)].(string)

	if !ok {
		return ""
	}
	return xuid
}

func SetCooldown(player *player.Player, key string, t int64, values []string) {
	if key == "combat" && !InCooldown(player, key) {
		player.Message("Vous êtes désormais en combat, vous ne pouvez plus vous téléporter ou vous déconnecter !")
	}

	user, err := session.GetUser(player)
	if err != nil {
		return
	}

	t = time.Now().Add(time.Duration(t) * time.Second).Unix()
	values = append([]string{strconv.Itoa(int(t))}, values...)

	user.Data.Cooldowns[key] = values
}

func InCooldown(player *player.Player, key string) bool {
	if key == "combat" && player.GameMode() == world.GameModeCreative {
		return false
	}

	user, err := session.GetUser(player)
	if err != nil {
		return false
	}

	cooldown, ok := user.Data.Cooldowns[key]
	if !ok {
		return false
	}

	cooldownTime, err := strconv.Atoi(cooldown[0])
	if err != nil {
		return false
	}

	return time.Now().Unix() < int64(cooldownTime)
}

func GetCooldownData(player *player.Player, key string) []string {
	user, err := session.GetUser(player)
	if err != nil {
		return []string{}
	}

	cooldown, ok := user.Data.Cooldowns[key]
	if !ok {
		return []string{"0", ""}
	}

	return cooldown
}

func GetCooldownTime(player *player.Player, key string) int {
	data := GetCooldownData(player, key)

	cooldownTime, err := strconv.Atoi(data[0])
	if err != nil {
		return 0
	}

	return cooldownTime - int(time.Now().Unix())
}

func RemoveCooldown(player *player.Player, key string) {
	user, err := session.GetUser(player)
	if err != nil {
		return
	}

	delete(user.Data.Cooldowns, key)
}

func InsideZone(position mgl64.Vec3, zone string) bool {
	coords := strings.Split(util.Zones[zone], ":")

	x1, _ := strconv.Atoi(coords[0])
	y1, _ := strconv.Atoi(coords[1])
	z1, _ := strconv.Atoi(coords[2])

	x2, _ := strconv.Atoi(coords[3])
	y2, _ := strconv.Atoi(coords[4])
	z2, _ := strconv.Atoi(coords[5])

	minX := math.Min(float64(x1), float64(x2))
	minY := math.Min(float64(y1), float64(y2))
	minZ := math.Min(float64(z1), float64(z2))

	maxX := math.Max(float64(x1), float64(x2))
	maxY := math.Max(float64(y1), float64(y2))
	maxZ := math.Max(float64(z1), float64(z2))

	x := position.X()
	y := position.Y()
	z := position.Z()

	return x >= minX && x <= maxX && y >= minY && y <= maxY && z >= minZ && z <= maxZ
}
