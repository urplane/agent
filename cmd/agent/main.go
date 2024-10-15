package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/urplane/agent"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	log.Info("configuring the agent", slog.String("version", agent.Version))

	data := make(map[string]json.RawMessage, 0)
	dataMu := &sync.Mutex{}

	var wg sync.WaitGroup

	for typ, fetcher := range agent.Fetchers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			log.Info("fetching " + typ + " data")
			if err := fetcher.Fetch(ctx); err != nil {
				log.Error("could not fetch "+typ+" data", slog.Any("err", err))
				return
			}

			b, err := fetcher.MarshalData()
			if err != nil {
				log.Error("could not marshal "+typ+" data", slog.Any("err", err))
				return
			}

			dataMu.Lock()
			data[typ] = b
			dataMu.Unlock()
		}()
	}

	wg.Wait()

	// TODO: Send the data to Urplane's API endpoint.
}
