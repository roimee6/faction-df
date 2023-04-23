package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"math/rand"
	"strings"
	"time"
)

type Bienvenue struct{}

func (Bienvenue) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	if handler.LastJoin == "" || util.InArray(sender.Name(), handler.AlreadyWished) {
		sender.Message("Vous avez déjà souhaité la bienvenue ou aucun nouveau joueur n'a rejoint le serveur dernièrement")
		return
	}

	rand.Seed(time.Now().UnixNano())

	message := util.WelcomeMessages[rand.Intn(len(util.WelcomeMessages))]
	message = handler.GetMessage(sender, strings.Replace(message, "{player}", handler.LastJoin, -1))

	target, ok := util.Server.PlayerByName(handler.LastJoin)

	if ok && target.Name() != sender.Name() {
		target.Message(message)
	}

	handler.AlreadyWished = append(handler.AlreadyWished, sender.Name())
	sender.Message(message)

	user.Data.Money += 500
	sender.Message("Vous venez de recevoir 500 pièces pour avoir souhaité la bienvenue à " + target.Name())
}
