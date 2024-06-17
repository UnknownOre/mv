package mv622

import (
	"github.com/oomph-ac/mv/multiversion/mv622/packet"
	"github.com/oomph-ac/mv/multiversion/mv630"
	"github.com/oomph-ac/mv/multiversion/util"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"

	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Protocol struct{}

func (Protocol) ID() int32 {
	return 622
}

func (Protocol) Ver() string {
	return "1.20.40"
}

func (Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReader(r, shieldID, enableLimits)
}

func (Protocol) NewWriter(w minecraft.ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriter(w, shieldID)
}

func (Protocol) Packets(listener bool) gtpacket.Pool {
	if listener {
		return packet.NewClientPool()
	}
	return packet.NewServerPool()
}

func (Protocol) ConvertToLatest(pk gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	if upgraded, ok := util.DefaultUpgrade(conn, pk, Mapping); ok {
		if upgraded == nil {
			return []gtpacket.Packet{}
		}

		return Upgrade([]gtpacket.Packet{upgraded}, conn)
	}

	return Upgrade([]gtpacket.Packet{pk}, conn)
}

func (Protocol) ConvertFromLatest(pk gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	if downgraded, ok := util.DefaultDowngrade(conn, pk, Mapping); ok {
		return Downgrade([]gtpacket.Packet{downgraded}, conn)
	}

	return Downgrade([]gtpacket.Packet{pk}, conn)
}

func Upgrade(pks []gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	return mv630.Upgrade(pks, conn)
}

func Downgrade(pks []gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	packets := make([]gtpacket.Packet, 0, len(pks))
	for _, pk := range mv630.Downgrade(pks, conn) {
		switch pk := pk.(type) {
		case *gtpacket.ShowStoreOffer:
			packets = append(packets, &packet.ShowStoreOffer{
				OfferID: pk.OfferID,
				ShowAll: false, // I don't think we can really translate this one.
			})
		case *gtpacket.SetPlayerInventoryOptions, *gtpacket.PlayerToggleCrafterSlotRequest:
			// These packets are not supported in 1.20.40, so we just ignore them.
			continue
		default:
			packets = append(packets, pk)
		}
	}

	pks = nil
	return packets
}
