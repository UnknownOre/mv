package mv649

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"

	"github.com/oomph-ac/mv/multiversion/mv649/packet"
	"github.com/oomph-ac/mv/multiversion/util"
	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Protocol struct{}

func (Protocol) ID() int32 {
	return 649
}

func (Protocol) Ver() string {
	return "1.20.60"
}

func (Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReader(r, shieldID, enableLimits)
}

func (Protocol) NewWriter(r minecraft.ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriter(r, shieldID)
}

func (Protocol) Packets(listener bool) gtpacket.Pool {
	if listener {
		return packet.NewClientPool()
	}
	return packet.NewServerPool()
}

func (Protocol) Encryption(key [32]byte) gtpacket.Encryption {
	return gtpacket.NewCTREncryption(key[:])
}

func (Protocol) ConvertToLatest(pk gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	fmt.Printf("upgrade %T\n", pk)

	if upgraded, ok := util.DefaultUpgrade(conn, pk, Mapping); ok {
		if upgraded == nil {
			return []gtpacket.Packet{}
		}

		return Upgrade([]gtpacket.Packet{upgraded}, conn)
	}

	return Upgrade([]gtpacket.Packet{pk}, conn)
}

func (Protocol) ConvertFromLatest(pk gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	fmt.Printf("downgrade %T\n", pk)

	if downgraded, ok := util.DefaultDowngrade(conn, pk, Mapping); ok {
		if downgraded == nil {
			return []gtpacket.Packet{}
		}

		return Downgrade([]gtpacket.Packet{downgraded}, conn)
	}

	return Downgrade([]gtpacket.Packet{pk}, conn)
}

func Upgrade(pks []gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	packets := []gtpacket.Packet{}
	for _, pk := range pks {
		switch pk := pk.(type) {
		case *packet.PlayerAuthInput:
			packets = append(packets, &gtpacket.PlayerAuthInput{
				Pitch:                  pk.Pitch,
				Yaw:                    pk.Yaw,
				Position:               pk.Position,
				MoveVector:             pk.MoveVector,
				HeadYaw:                pk.HeadYaw,
				InputData:              pk.InputData,
				InputMode:              pk.InputMode,
				PlayMode:               pk.PlayMode,
				InteractionModel:       pk.InteractionModel,
				GazeDirection:          pk.GazeDirection,
				Tick:                   pk.Tick,
				Delta:                  pk.Delta,
				ItemInteractionData:    pk.ItemInteractionData,
				ItemStackRequest:       pk.ItemStackRequest,
				BlockActions:           pk.BlockActions,
				ClientPredictedVehicle: pk.ClientPredictedVehicle,
				AnalogueMoveVector:     pk.AnalogueMoveVector,
				VehicleRotation:        mgl32.Vec2{},
			})
		case *gtpacket.LecternUpdate:
			packets = append(packets, &packet.LecternUpdate{
				Page:      pk.Page,
				PageCount: pk.PageCount,
				Position:  pk.Position,
				DropBook:  false,
			})
		default:
			packets = append(packets, pk)
		}
	}

	return packets
}

func Downgrade(pks []gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	packets := []gtpacket.Packet{}
	for _, pk := range pks {
		switch pk := pk.(type) {
		case *gtpacket.AvailableCommands:
			for _, c := range pk.Commands {
				for _, o := range c.Overloads {
					for _, p := range o.Parameters {
						switch p.Type {
						case protocol.CommandArgTypeEquipmentSlots:
							p.Type = packet.CommandArgTypeEquipmentSlots
						case protocol.CommandArgTypeString:
							p.Type = packet.CommandArgTypeString
						case protocol.CommandArgTypeBlockPosition:
							p.Type = packet.CommandArgTypeBlockPosition
						case protocol.CommandArgTypePosition:
							p.Type = packet.CommandArgTypePosition
						case protocol.CommandArgTypeMessage:
							p.Type = packet.CommandArgTypeMessage
						case protocol.CommandArgTypeRawText:
							p.Type = packet.CommandArgTypeRawText
						case protocol.CommandArgTypeJSON:
							p.Type = packet.CommandArgTypeJSON
						case protocol.CommandArgTypeBlockStates:
							p.Type = packet.CommandArgTypeBlockStates
						case protocol.CommandArgTypeCommand:
							p.Type = packet.CommandArgTypeCommand
						}
					}
				}
			}

			packets = append(packets, pk)
		case *gtpacket.SetActorMotion:
			packets = append(packets, &packet.SetActorMotion{
				Velocity:        pk.Velocity,
				EntityRuntimeID: pk.EntityRuntimeID,
			})
		case *gtpacket.ResourcePacksInfo:
			packets = append(packets, &packet.ResourcePacksInfo{
				TexturePackRequired: pk.TexturePackRequired,
				HasScripts:          pk.HasScripts,
				BehaviourPacks:      pk.BehaviourPacks,
				TexturePacks:        pk.TexturePacks,
				ForcingServerPacks:  pk.ForcingServerPacks,
				PackURLs:            pk.PackURLs,
			})
		case *gtpacket.MobEffect:
			packets = append(packets, &packet.MobEffect{
				EntityRuntimeID: pk.EntityRuntimeID,
				Operation:       pk.Operation,
				EffectType:      pk.EffectType,
				Amplifier:       pk.Amplifier,
				Particles:       pk.Particles,
				Duration:        pk.Duration,
			})
		default:
			packets = append(packets, pk)
		}
	}

	return packets
}
