package session

import (
	"encoding/json"
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/util"
	"os"
)

type Data struct {
	Kill       int
	Death      int
	Money      int
	Killstreak int

	Coordinates  bool
	FactionChat  bool
	StaffChat    bool
	PlayerBefore bool
	Scoreboard   bool

	Faction *string

	Rank        string
	Name        string
	DisplayName string

	Cooldowns map[string][]string

	Invites       []string
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
		SaveData(p.XUID(), user.Data)
	}

	delete(users, p.Name())
}

func loadUserData(p *player.Player) Data {
	data := ParseData(p.XUID())

	fmt.Println(data)

	if data.Rank == "" {
		data.Rank = "joueur"
		data.Money = 1000
		data.DisplayName = p.Name()
		data.FactionChat = false
		data.StaffChat = false
		data.PlayerBefore = false
		data.Scoreboard = true
	}

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

	if data.Coordinates {
		p.ShowCoordinates()
	}
	return data
}

func SaveData(xuid string, data Data) {
	filename := "asset/data/players/" + xuid + ".json"

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

	if jsonBytes, err := json.Marshal(data); err == nil {
		_ = os.WriteFile(filename, jsonBytes, 0644)
	}
}

func ParseData(xuid string) Data {
	file, err := os.ReadFile("asset/data/players/" + xuid + ".json")
	data := Data{}

	data.Cooldowns = make(map[string][]string, 0)

	if err == nil {
		_ = json.Unmarshal(file, &data)
	}
	return data
}
