package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	nitro "github.com/roimee6/Faction/server"
	"github.com/roimee6/Faction/server/handler"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	log := logrus.New()

	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		handler.SaveCache()
	}()

	srv := conf.New()

	srv.CloseOnProgramEnd()
	srv.Listen()

	nitro.New(srv)

	for srv.Accept(func(p *player.Player) {
		h := nitro.NewHandler(p)
		p.Handle(h)

		nitro.HandleJoin(p)
	}) {
	}
}

func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config

	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)

		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}

		return c.Config(log)
	}

	data, err := os.ReadFile("config.toml")

	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}

	c.Server.JoinMessage = ""
	c.Server.QuitMessage = ""

	return c.Config(log)
}
