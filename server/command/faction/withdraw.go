package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"math"
	"strconv"
)

type Withdraw struct {
	Sub    cmd.SubCommand `cmd:"withdraw"`
	Amount int            `cmd:"amount"`
}

func (w Withdraw) Run(source cmd.Source, _ *cmd.Output) {
	sender, ok := source.(*player.Player)
	if !ok {
		return
	}

	if !handler.HasFaction(sender) {
		sender.Message("Vous n'avez pas de faction !")
		return
	}

	user, err := session.GetUser(sender)
	if err != nil {
		return
	}

	faction := user.Data.Faction
	amount := int(math.Floor(float64(w.Amount)))

	if 0 >= amount {
		sender.Message("Vous ne pouvez pas retirer une somme négative ou nulle !")
		return
	} else if amount >= handler.Factions[*faction].Money {
		sender.Message("Vous n'avez pas assez d'argent !")
		return
	}

	handler.Factions[*faction].Money -= amount
	user.Data.Money += amount

	sender.Messagef("Vous venez de retirer %s de la faction !", strconv.Itoa(amount))
}
