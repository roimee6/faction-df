package session

import (
	"encoding/json"
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/util"
	"io/ioutil"
	"os"
)

type Data struct {
	Kill  int
	Death int

	Coordinates bool

	Faction *string
	Name    string

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

var users = make(map[string]*User)

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
	if _, ok := users[p.Name()]; !ok {
		return nil, fmt.Errorf("user %s not found", p.Name())
	}
	return users[p.Name()], nil
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
