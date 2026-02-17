package unsafe

import (
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Session(p *player.Player) *session.Session {
	return player_session(p)
}

func Conn(p *player.Player) session.Conn {
	s := player_session(p)
	return fetchPrivateField[session.Conn](s, "conn")
}

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

func playerEntityData(p *player.Player) *world.EntityData {
	return fetchPrivateField[*world.EntityData](p, "data")
}

func SetPlayerMovementGravity(p *player.Player, gravity float64) {
	if p == nil {
		return
	}
	mc := fetchPrivateField[*entity.MovementComputer](p, "mc")
	if mc == nil {
		return
	}
	mc.Gravity = gravity
}

func AddPlayerRotation(p *player.Player, dyaw, dpitch float64) cube.Rotation {
	if p == nil {
		return cube.Rotation{}
	}
	data := playerEntityData(p)
	rot := data.Rot.Add(cube.Rotation{dyaw, dpitch})
	data.Rot = rot
	return rot
}

func SetPlayerRotation(p *player.Player, yaw, pitch float64) {
	if p == nil {
		return
	}
	data := playerEntityData(p)
	data.Rot = cube.Rotation{yaw, pitch}

	for _, v := range p.Tx().Viewers(p.Position()) {
		v.ViewEntityMovement(p, p.Position(), cube.Rotation{yaw, pitch}, p.OnGround())
	}
}

func SetHeldSlot(p *player.Player, slot int) {
	if p == nil {
		return
	}
	updatePrivateField(p, "heldSlot", slot)

	for _, v := range p.Tx().Viewers(p.Position()) {
		v.ViewEntityItems(p)
	}
}
