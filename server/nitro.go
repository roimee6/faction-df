package server

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/roimee6/Faction/server/command"
	"github.com/roimee6/Faction/server/command/faction"
	"github.com/roimee6/Faction/server/util"
)

func New(srv *server.Server) {
	util.Server = srv

	util.LoadCache()

	loadCommands()
}

func loadCommands() {
	cmd.Register(cmd.New("xyz", "Désactive ou active les coordonnées", nil, command.Xyz{}))
	cmd.Register(cmd.New("gamemode", "Change de mode de jeu", []string{"gm"}, command.GameModeString{}, command.GameModeInt{}))
	cmd.Register(cmd.New("weather", "Change le temps", nil, command.Weather{}))
	cmd.Register(cmd.New("time", "Change l'heure", nil, command.SetTimeInt{}, command.SetTimeString{}, command.Add{}))
	cmd.Register(cmd.New("teleport", "Téléporte un joueur", []string{"tp"}, command.TeleportXYZ{}, command.TeleportPlayer{}))
	cmd.Register(cmd.New("stop", "Arrête le serveur", nil, command.Stop{}))
	cmd.Register(cmd.New("say", "Envoie un message dans le chat", nil, command.Say{}))
	cmd.Register(cmd.New("sudo", "Execute une commande en tant qu'un autre joueur", nil, command.Sudo{}))

	cmd.Register(cmd.New("money", "regarde la money", nil, command.Money{}))

	cmd.Register(cmd.New("faction", "Commande de faction", []string{"f"}, faction.Info{}, faction.Accept{}, faction.Create{}, faction.Delete{}, faction.Delhome{}, faction.Deposit{}, faction.Home{}, faction.Invite{}, faction.Leave{}, faction.Sethome{}, faction.Withdraw{}))
}
