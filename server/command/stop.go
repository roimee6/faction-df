package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/util"
)

type Stop struct{}

func (t Stop) Run(_ cmd.Source, output *cmd.Output) {
	output.Printf("Stopping server.")

	err := util.Server.Close()

	if err != nil {
		output.Printf("error shutting down server: %v", err)
	}

	panic("stop")
}

func (Stop) Allow(source cmd.Source) bool {
	return handler.HasRankPermission(source, "fondateur")
}
