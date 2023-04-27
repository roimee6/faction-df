package command

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"strings"
)

type Message struct {
	Targets []cmd.Target `cmd:"target"`
	Message cmd.Varargs  `cmd:"message"`
}

func (m Message) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if len(m.Targets) < 1 {
		sender.Message("Aucune cible spécifiée !")
		return
	}

	target, ok := m.Targets[0].(*player.Player)

	if !ok {
		sender.Message("Le joueur indiqué n'est pas connecté sur le serveur")
		return
	} else if handler.InCooldown(sender, "mute") {
		sender.Message("Vous etes actuellement mute, temps restant: " + util.FormatSeconds(handler.GetCooldownTime(sender, "mute"), 0))
		return
	}

	targetUsername := target.Name()
	user.Data.LastReply = &targetUsername

	targetUser, err := session.GetUser(target)
	if err != nil {
		return
	}

	senderUsername := sender.Name()
	targetUser.Data.LastReply = &senderUsername

	message := strings.TrimSpace(string(m.Message))

	fmt.Println("[MP] [" + senderUsername + " » " + targetUsername + "] " + message)

	sender.Message("[MP] [" + senderUsername + " » " + targetUsername + "] " + message)
	target.Message("[MP] [" + senderUsername + " » " + targetUsername + "] " + message)

	sender.PlaySound(sound.Click{})
	target.PlaySound(sound.Click{})
}
