package agent

import (
	"context"
	"encoding/json"
)

type NetworkFetcher struct {
	data NetworkData
}

var _ Fetcher = (*NetworkFetcher)(nil)

func NewNetworkFetcher() *NetworkFetcher {
	return &NetworkFetcher{}
}

func (n *NetworkFetcher) Fetch(ctx context.Context) error {
	return nil
}

func (n *NetworkFetcher) MarshalData() ([]byte, error) {
	return json.Marshal(n.data)
}

// TODO: Complete this.
type NetworkData struct{}
