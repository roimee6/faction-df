package server

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/roimee6/Faction/server/handler"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"math"
	"strconv"
	"strings"
	"time"
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

func (h *Handler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()

	user, err := session.GetUser(h.p)
	if err != nil {
		return
	}

	if !handler.HasRankPermission(h.p, "chat") {
		if handler.InCooldown(h.p, "chat") {
			return
		}

		handler.SetCooldown(h.p, "chat", 1, []string{})
	}

	if (user.Data.FactionChat || (*message)[0:1] == "-") && handler.HasFaction(h.p) {
		faction := *user.Data.Faction

		if (*message)[0:1] == "-" {
			*message = (*message)[1:]
		}

		fmt.Println("[F] [" + faction + "] " + h.p.Name() + " » " + (*message) + "\n")
		handler.BroadcastFactionMessage(faction, "[F] "+h.p.Name()+" » "+(*message))

		return
	} else if (user.Data.StaffChat || (*message)[0:1] == "!") && handler.HasRankPermission(h.p, "guide") {
		if (*message)[0:1] == "!" {
			*message = (*message)[1:]
		}

		fmt.Printf("[StaffChat] [" + *user.Data.Faction + "] " + h.p.Name() + " » " + (*message) + "\n")
		handler.BroadcastFactionMessage(*user.Data.Faction, "[StaffChat] "+h.p.Name()+" » "+(*message))

		return
	} else if handler.InCooldown(h.p, "mute") {
		h.p.Message("Vous etes actuellement mute, temps restant: " + util.FormatSeconds(handler.GetCooldownTime(h.p, "mute"), 0))
		return
	}

	_, _ = chat.Global.WriteString(handler.GetMessage(h.p, *message))
}

func (h *Handler) HandleCommandExecution(_ *event.Context, command cmd.Command, args []string) {
	fmt.Printf("[%s] /%s %s\n", h.p.Name(), command.Name(), strings.Join(args, " "))
}

func (h *Handler) HandleDeath(source world.DamageSource, _ *bool) {
	s, ok := source.(entity.AttackDamageSource)

	user, err := session.GetUser(h.p)
	if err != nil {
		return
	}

	handler.RemoveCooldown(h.p, "combat")

	user.Data.Death += 1
	user.Data.Killstreak = 0

	if handler.HasFaction(h.p) {
		handler.Factions[*user.Data.Faction].Power -= 4
	}

	if ok {
		if d, ok := s.Attacker.(*player.Player); ok {
			userD, err := session.GetUser(d)
			if err != nil {
				return
			}

			userD.Data.Kill += 1
			userD.Data.Killstreak += 1

			if handler.HasFaction(d) {
				handler.Factions[*userD.Data.Faction].Power += 6
			}

			if userD.Data.Killstreak%5 == 0 {
				_, _ = chat.Global.WriteString("§c" + d.Name() + " est en killstreak de " + strconv.Itoa(userD.Data.Killstreak) + "§c !")
			}
		}
	}
}

func HandleJoin(p *player.Player) {
	session.CreateUser(p)

	user, err := session.GetUser(p)
	if err != nil {
		return
	}

	handler.UpdateNameTag(p)
	handler.PushData(p.XUID(), p.Name(), user.Data)

	if !user.Data.PlayerBefore {
		count := strconv.Itoa(len(handler.Players["xuid"]) + 1)
		user.Data.PlayerBefore = true

		_, _ = chat.Global.WriteString(p.Name() + " a rejoint le serveur pour la première fois souhaitez lui la bienvenue avec la commande /bvn (#" + count + ") !")

		handler.LastJoin = p.Name()
		handler.AlreadyWished = []string{}
	}

	util.SendTip("§a+ " + p.Name() + " +")
}

func (h *Handler) HandleQuit() {
	session.RemoveUser(h.p)
	util.SendTip("§c- " + h.p.Name() + " -")
}

func (h *Handler) HandleHurt(ctx *event.Context, _ *float64, _ *time.Duration, src world.DamageSource) {
	var attacker *player.Player

	if _, ok := src.(entity.FallDamageSource); ok {
		ctx.Cancel()
		return

		// TODO ENLEVER OU NON ? CAR FAUT ACTIVER LES FALL DAMAGE
	}

	if handler.InsideZone(h.p.Position(), "spawn") {
		ctx.Cancel()
		return
	}

	if src, ok := src.(entity.AttackDamageSource); ok {
		if p, ok := src.Attacker.(*player.Player); ok {
			attacker = p
		}

		pUser, err := session.GetUser(h.p)
		if err != nil {
			return
		}

		aUser, err := session.GetUser(attacker)
		if err != nil {
			return
		}

		if handler.HasFaction(attacker) && handler.HasFaction(h.p) {
			if *pUser.Data.Faction == *aUser.Data.Faction {
				ctx.Cancel()
				return
			}
		}

		if attacker.GameMode() == world.GameModeCreative || h.p.GameMode() == world.GameModeCreative {
			return
		}

		handler.SetCooldown(h.p, "combat", 30, []string{attacker.Name()})
		handler.SetCooldown(attacker, "combat", 30, []string{h.p.Name()})

		health := strconv.Itoa(int(math.Round(h.p.Health()*100) / 100))
		h.p.SetScoreTag(health + "§c❤")
	}
}

func (h *Handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, _ *bool) {
	*force, *height = 0.383, 0.387

	if o, ok := e.(interface{ OnGround() bool }); ok && !o.OnGround() {
		if dist := e.Position().Y() - h.p.Position().Y(); dist >= 3 {
			*height -= dist / 28.795
		}
	}

	if handler.InsideZone(h.p.Position(), "spawn") {
		ctx.Cancel()
		return
	}
}
