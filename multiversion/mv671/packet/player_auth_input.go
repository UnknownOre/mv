package packet

const (
	InputFlagAscend = 1 << iota
	InputFlagDescend
	InputFlagNorthJump
	InputFlagJumpDown
	InputFlagSprintDown
	InputFlagChangeHeight
	InputFlagJumping
	InputFlagAutoJumpingInWater
	InputFlagSneaking
	InputFlagSneakDown
	InputFlagUp
	InputFlagDown
	InputFlagLeft
	InputFlagRight
	InputFlagUpLeft
	InputFlagUpRight
	InputFlagWantUp
	InputFlagWantDown
	InputFlagWantDownSlow
	InputFlagWantUpSlow
	InputFlagSprinting
	InputFlagAscendBlock
	InputFlagDescendBlock
	InputFlagSneakToggleDown
	InputFlagPersistSneak
	InputFlagStartSprinting
	InputFlagStopSprinting
	InputFlagStartSneaking
	InputFlagStopSneaking
	InputFlagStartSwimming
	InputFlagStopSwimming
	InputFlagStartJumping
	InputFlagStartGliding
	InputFlagStopGliding
	InputFlagPerformItemInteraction
	InputFlagPerformBlockActions
	InputFlagPerformItemStackRequest
	InputFlagHandledTeleport
	InputFlagEmoting
	InputFlagMissedSwing
	InputFlagStartCrawling
	InputFlagStopCrawling
	InputFlagStartFlying
	InputFlagStopFlying
	InputFlagClientAckServerData
	InputFlagClientPredictedVehicle
	InputFlagPaddlingLeft
	InputFlagPaddlingRight
)

const (
	InputModeMouse = iota + 1
	InputModeTouch
	InputModeGamePad
	InputModeMotionController
)

const (
	PlayModeNormal = iota
	PlayModeTeaser
	PlayModeScreen
	PlayModeViewer
	PlayModeReality
	PlayModePlacement
	PlayModeLivingRoom
	PlayModeExitLevel
	PlayModeExitLevelLivingRoom
	PlayModeNumModes
)

const (
	InteractionModelTouch = iota
	InteractionModelCrosshair
	InteractionModelClassic
)
