package unsafe

import (
	_ "unsafe"

	"github.com/df-mc/atomic"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Session returns the underlying network session of the given player.
func Session(p *player.Player) *session.Session {
	return player_session(p)
}

// WritePacket writes the given packet to a network session or a player.
func WritePacket[T *player.Player | *session.Session](target T, pk packet.Packet) {
	s, ok := any(target).(*session.Session)
	if !ok {
		s = Session(any(target).(*player.Player))
	}

	if s == session.Nop {
		return
	}
	session_writePacket(s, pk)
}

// Rotate rotates the player with the given yaw and pitch.
func Rotate(p *player.Player, yaw, pitch float64) {
	updatePrivateField(p, "yaw", *atomic.NewFloat64(yaw))
	updatePrivateField(p, "pitch", *atomic.NewFloat64(pitch))

	for _, v := range p.World().Viewers(p.Position()) {
		v.ViewEntityMovement(p, p.Position(), cube.Rotation{yaw, pitch}, p.OnGround())
	}
}

// UpdateHeldSlot updates the held slot of the player.
func UpdateHeldSlot(p *player.Player, slot int) {
	updatePrivateField(p, "heldSlot", slot)
	
	for _, v := range p.World().Viewers(p.Position()) {
		v.ViewEntityItems(p)
	}
}
