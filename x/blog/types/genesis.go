package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PostCount: PostCount{
			Count: uint64(0),
		},
		StoredPostList: []StoredPost{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedPost
	storedPostIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredPostList {
		index := string(StoredPostKey(elem.Index))
		if _, ok := storedPostIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedPost")
		}
		storedPostIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
