package staking

import (
	"github.com/Ankr-network/ankr-protocol/shared/types"
)

type SortedDelegators []*types.Delegator

func (s SortedDelegators) Len() int {
	return len(s)
}

func (s SortedDelegators) Less(i, j int) bool {
	return s[i].Epoch > s[j].Epoch
}

func (s SortedDelegators) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
