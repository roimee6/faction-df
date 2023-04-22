package session

import (
	"encoding/json"
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/util"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Data struct {
	Kill  int
	Death int
	Money int

	Coordinates bool

	Faction *string
	Name    string

	Cooldowns map[string][]string

	Invites []string

	Addresses     []string
	UUIDs         []string
	SelfSignedIDs []string
	DeviceIDs     []string
}

type User struct {
	Player *player.Player
	Data   Data
}

var users = make(map[string]*User, 0)

func CreateUser(p *player.Player) User {
	if _, ok := users[p.Name()]; ok {
		return *users[p.Name()]
	}

	u := User{
		Player: p,
		Data:   loadUserData(p),
	}

	users[p.Name()] = &u
	return u
}

func GetUser(p *player.Player) (*User, error) {
	if u, ok := users[p.Name()]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("user %s not found", p.Name())
}

func RemoveUser(p *player.Player) {
	if user, err := GetUser(p); err == nil {
		saveUserData(user)
	}

	delete(users, p.Name())
}

func loadUserData(p *player.Player) Data {
	file, err := ioutil.ReadFile("asset/data/players/" + p.XUID() + ".json")
	data := Data{}

	data.Cooldowns = make(map[string][]string, 0)

	data.Name = p.Name()
	data.Invites = []string{}

	if !util.InArray(p.Addr().String(), data.Addresses) {
		data.Addresses = append(data.Addresses, p.Addr().String())
	}
	if !util.InArray(p.UUID().String(), data.UUIDs) {
		data.UUIDs = append(data.UUIDs, p.UUID().String())
	}
	if !util.InArray(p.SelfSignedID(), data.SelfSignedIDs) {
		data.SelfSignedIDs = append(data.SelfSignedIDs, p.SelfSignedID())
	}
	if !util.InArray(p.DeviceID(), data.DeviceIDs) {
		data.DeviceIDs = append(data.DeviceIDs, p.DeviceID())
	}

	if err != nil {
		_ = json.Unmarshal(file, &data)
	}

	if data.Coordinates {
		p.ShowCoordinates()
	}
	return data
}

func saveUserData(user *User) {
	filename := "asset/data/players/" + user.Player.XUID() + ".json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, _ := os.Create(filename)

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	file, _ := os.OpenFile(filename, os.O_WRONLY, 0644)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if jsonBytes, err := json.Marshal(user.Data); err == nil {
		_ = ioutil.WriteFile(filename, jsonBytes, 0644)
	}
}

func SetCooldown(player *player.Player, key string, t int64, values []string) {
	user, err := GetUser(player)
	if err != nil {
		return
	}

	t = time.Now().Add(time.Duration(t) * time.Second).Unix()
	values = append([]string{strconv.Itoa(int(t))}, values...)

	user.Data.Cooldowns[key] = values
}

func InCooldown(player *player.Player, key string) bool {
	user, err := GetUser(player)
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
	user, err := GetUser(player)
	if err != nil {
		return []string{}
	}

	cooldown, ok := user.Data.Cooldowns[key]
	if !ok {
		return []string{"0", ""}
	}

	return cooldown
}

func RemoveCooldown(player *player.Player, key string) {
	user, err := GetUser(player)
	if err != nil {
		return
	}

	delete(user.Data.Cooldowns, key)
}
