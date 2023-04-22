package faction

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"math"
	"strconv"
)

type Deposit struct {
	Sub    cmd.SubCommand `cmd:"deposit"`
	Amount int            `cmd:"amount"`
}

func (d Deposit) Run(source cmd.Source, _ *cmd.Output) {
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
	amount := int(math.Floor(float64(d.Amount)))

	if 0 >= amount {
		sender.Message("Vous ne pouvez pas donner une somme négative ou nulle !")
		return
	} else if amount >= user.Data.Money {
		sender.Message("Vous n'avez pas assez d'argent !")
		return
	}

	util.Factions[*faction].Money += amount
	user.Data.Money -= amount

	sender.Messagef("Vous venez de donner %s à la faction !", strconv.Itoa(amount))
}
