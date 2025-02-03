package pkgs

// Process Name Constants
// process : identifier
const (
	PersistState                  = "PersistState"
	FetchAllSlots                 = "FetchAllSlots"
	MonitorEvents                 = "MonitorEvents"
	ProcessSnapshotterStateEvents = "ProcessSnapshotterStateEvents"
	ProcessProtocolStateEvents    = "ProcessProtocolStateEvents"
)

// State Variable Constants
// state variable : identifier
const (
	CurrentEpochID             = "CurrentEpochID"
	CurrentDayKey              = "CurrentDayKey"
	TotalNodesCount            = "TotalNodesCount"
	EpochsInADay               = "EpochsInADay"
	EPOCH_SIZE                 = "EPOCH_SIZE"
	SOURCE_CHAIN_BLOCK_TIME    = "SOURCE_CHAIN_BLOCK_TIME"
	AllSlotsInfoKey            = "AllSlotsInfoKey"
	DailySnapshotQuotaTableKey = "DailySnapshotQuotaTableKey"
)
