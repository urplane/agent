package agent

import "context"

type Fetcher interface {
	Fetch(ctx context.Context) error
	MarshalData() ([]byte, error)
}

var Fetchers = map[string]Fetcher{
	"network": NewNetworkFetcher(),
}
