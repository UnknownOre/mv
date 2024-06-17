package packet

import (
	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// NOTE: CorrectPlayerMovementPrediction is not included in here, since changes
// to the packet were made late, and it was updated around 1.20.50 (630).

const (
	IDText              uint32 = 9
	IDStartGame         uint32 = 11
	IDContainerClose    uint32 = 47
	IDCodeBuilderSource uint32 = 150
)

func NewClientPool() gtpacket.Pool {
	pool := gtpacket.NewClientPool()
	pool[IDText] = func() gtpacket.Packet { return &Text{} }
	pool[IDCodeBuilderSource] = func() gtpacket.Packet { return &CodeBuilderSource{} }
	pool[IDContainerClose] = func() gtpacket.Packet { return &ContainerClose{} }

	return pool
}

func NewServerPool() gtpacket.Pool {
	pool := gtpacket.NewServerPool()
	pool[IDText] = func() gtpacket.Packet { return &Text{} }
	pool[IDStartGame] = func() gtpacket.Packet { return &StartGame{} }
	pool[IDContainerClose] = func() gtpacket.Packet { return &ContainerClose{} }
	pool[IDCodeBuilderSource] = func() gtpacket.Packet { return &CodeBuilderSource{} }

	return pool
}
