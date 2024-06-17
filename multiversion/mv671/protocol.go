package mv671

import (
	"github.com/oomph-ac/mv/multiversion/mv671/packet"
	"github.com/oomph-ac/mv/multiversion/util"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Protocol struct{}

func (Protocol) ID() int32 {
	return 671
}

func (Protocol) Ver() string {
	return "1.20.80"
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
	var packets []gtpacket.Packet
	for _, pk := range pks {
		switch pk := pk.(type) {
		case *packet.CodeBuilderSource:
			for _, v := range pk.Value {
				packets = append(packets, &gtpacket.CodeBuilderSource{
					Operation:  pk.Operation,
					Category:   pk.Category,
					CodeStatus: v,
				})
			}
		case *packet.Text:
			packets = append(packets, &gtpacket.Text{
				TextType:         pk.TextType,
				NeedsTranslation: pk.NeedsTranslation,
				SourceName:       pk.SourceName,
				Message:          pk.Message,
				Parameters:       pk.Parameters,
				XUID:             pk.XUID,
				PlatformChatID:   pk.PlatformChatID,
				FilteredMessage:  pk.Message,
			})
		case *packet.ContainerClose:
			packets = append(packets, &gtpacket.ContainerClose{
				WindowID:      pk.WindowID,
				ContainerType: 0,
				ServerSide:    pk.ServerSide,
			})
		default:
			packets = append(packets, pk)
		}
	}

	return packets
}

func Downgrade(pks []gtpacket.Packet, conn *minecraft.Conn) []gtpacket.Packet {
	packets := make([]gtpacket.Packet, 0, len(pks))

	for _, pk := range pks {
		switch pk := pk.(type) {
		case *gtpacket.ContainerClose:
			packets = append(packets, &packet.ContainerClose{
				WindowID:   pk.WindowID,
				ServerSide: pk.ServerSide,
			})
		case *gtpacket.CodeBuilderSource:
			packets = append(packets, &packet.CodeBuilderSource{
				Operation: pk.Operation,
				Category:  pk.Category,
				Value: []byte{
					pk.CodeStatus,
				},
			})
		case *gtpacket.Text:
			packets = append(packets, &packet.Text{
				TextType:         pk.TextType,
				NeedsTranslation: pk.NeedsTranslation,
				SourceName:       pk.SourceName,
				Message:          pk.Message,
				Parameters:       pk.Parameters,
				XUID:             pk.XUID,
				PlatformChatID:   pk.PlatformChatID,
			})
		case *gtpacket.StartGame:
			packets = append(packets, &packet.StartGame{
				EntityUniqueID:                 pk.EntityUniqueID,
				EntityRuntimeID:                pk.EntityRuntimeID,
				PlayerGameMode:                 pk.PlayerGameMode,
				PlayerPosition:                 pk.PlayerPosition,
				Pitch:                          pk.Pitch,
				Yaw:                            pk.Yaw,
				WorldSeed:                      pk.WorldSeed,
				SpawnBiomeType:                 pk.SpawnBiomeType,
				UserDefinedBiomeName:           pk.UserDefinedBiomeName,
				Dimension:                      pk.Dimension,
				Generator:                      pk.Generator,
				WorldGameMode:                  pk.WorldGameMode,
				Difficulty:                     pk.Difficulty,
				WorldSpawn:                     pk.WorldSpawn,
				AchievementsDisabled:           pk.AchievementsDisabled,
				EditorWorldType:                pk.EditorWorldType,
				CreatedInEditor:                pk.CreatedInEditor,
				ExportedFromEditor:             pk.ExportedFromEditor,
				DayCycleLockTime:               pk.DayCycleLockTime,
				EducationEditionOffer:          pk.EducationEditionOffer,
				EducationFeaturesEnabled:       pk.EducationFeaturesEnabled,
				EducationProductID:             pk.EducationProductID,
				RainLevel:                      pk.RainLevel,
				LightningLevel:                 pk.LightningLevel,
				ConfirmedPlatformLockedContent: pk.ConfirmedPlatformLockedContent,
				MultiPlayerGame:                pk.MultiPlayerGame,
				LANBroadcastEnabled:            pk.LANBroadcastEnabled,
				XBLBroadcastMode:               pk.XBLBroadcastMode,
				PlatformBroadcastMode:          pk.PlatformBroadcastMode,
				CommandsEnabled:                pk.CommandsEnabled,
				TexturePackRequired:            pk.TexturePackRequired,
				GameRules:                      pk.GameRules,
				Experiments:                    pk.Experiments,
				ExperimentsPreviouslyToggled:   pk.ExperimentsPreviouslyToggled,
				BonusChestEnabled:              pk.BonusChestEnabled,
				StartWithMapEnabled:            pk.StartWithMapEnabled,
				PlayerPermissions:              pk.PlayerPermissions,
				ServerChunkTickRadius:          pk.ServerChunkTickRadius,
				HasLockedBehaviourPack:         pk.HasLockedBehaviourPack,
				HasLockedTexturePack:           pk.HasLockedTexturePack,
				FromLockedWorldTemplate:        pk.FromLockedWorldTemplate,
				MSAGamerTagsOnly:               pk.MSAGamerTagsOnly,
				FromWorldTemplate:              pk.FromWorldTemplate,
				WorldTemplateSettingsLocked:    pk.WorldTemplateSettingsLocked,
				OnlySpawnV1Villagers:           pk.OnlySpawnV1Villagers,
				PersonaDisabled:                pk.PersonaDisabled,
				CustomSkinsDisabled:            pk.CustomSkinsDisabled,
				EmoteChatMuted:                 pk.EmoteChatMuted,
				BaseGameVersion:                pk.BaseGameVersion,
				LimitedWorldWidth:              pk.LimitedWorldWidth,
				LimitedWorldDepth:              pk.LimitedWorldDepth,
				NewNether:                      pk.NewNether,
				EducationSharedResourceURI:     pk.EducationSharedResourceURI,
				LevelID:                        pk.LevelID,
				WorldName:                      pk.WorldName,
				TemplateContentIdentity:        pk.TemplateContentIdentity,
				Trial:                          pk.Trial,
				PlayerMovementSettings:         pk.PlayerMovementSettings,
				Time:                           pk.Time,
				EnchantmentSeed:                pk.EnchantmentSeed,
				Blocks:                         pk.Blocks,
				Items:                          pk.Items,
				MultiPlayerCorrelationID:       pk.MultiPlayerCorrelationID,
				ServerAuthoritativeInventory:   pk.ServerAuthoritativeInventory,
				GameVersion:                    pk.GameVersion,
				PropertyData:                   pk.PropertyData,
				ServerBlockStateChecksum:       pk.ServerBlockStateChecksum,
				ClientSideGeneration:           pk.ClientSideGeneration,
				WorldTemplateID:                pk.WorldTemplateID,
				ChatRestrictionLevel:           pk.ChatRestrictionLevel,
				DisablePlayerInteractions:      pk.DisablePlayerInteractions,
				UseBlockNetworkIDHashes:        pk.UseBlockNetworkIDHashes,
				ServerAuthoritativeSound:       pk.ServerAuthoritativeSound,
			})
		case *gtpacket.CraftingData:
			recipies := make([]protocol.Recipe, 0, len(pk.Recipes))
			for _, r := range pk.Recipes {
				switch r := r.(type) {
				case *protocol.ShapedRecipe:
					recipies = append(recipies, &packet.ShapedRecipe{
						RecipeID:        r.RecipeID,
						Width:           r.Width,
						Height:          r.Height,
						Input:           r.Input,
						Output:          r.Output,
						UUID:            r.UUID,
						Block:           r.Block,
						Priority:        r.Priority,
						RecipeNetworkID: r.RecipeNetworkID,
					})
				case *protocol.ShapelessRecipe:
					recipies = append(recipies, &packet.ShapelessRecipe{
						RecipeID:        r.RecipeID,
						Input:           r.Input,
						Output:          r.Output,
						UUID:            r.UUID,
						Block:           r.Block,
						Priority:        r.Priority,
						RecipeNetworkID: r.RecipeNetworkID,
					})
				default:
					recipies = append(recipies, r)
				}
			}

			packets = append(packets, &packet.CraftingData{
				Recipes:                      recipies,
				PotionRecipes:                pk.PotionRecipes,
				PotionContainerChangeRecipes: pk.PotionContainerChangeRecipes,
				MaterialReducers:             pk.MaterialReducers,
				ClearRecipes:                 pk.ClearRecipes,
			})
		default:
			packets = append(packets, pk)
		}
	}

	return packets
}
