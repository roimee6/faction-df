package session

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type Handler struct {
	player.NopHandler
	p *player.Player
}

func NewHandler(p *player.Player) *Handler {
	return &Handler{
		p: p,
	}
}

func (h *Handler) HandleChat(_ *event.Context, text *string) {
	fmt.Println(text)
}

func (h *Handler) HandleQuit() {
	RemoveUser(h.p)
}
