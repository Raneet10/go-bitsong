package types

const (
	// ModuleName is the name of the module
	ModuleName = "reward"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

// Keys for reward store
// Items are stored with the following key: values
//
// - 0x00: RewardPool
var (
	RewardPoolKey = []byte{0x00}
)