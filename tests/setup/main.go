package main

import (
	"os"

	statsig "github.com/statsig-io/go-sdk"
)

func main() {
	println("Preparing Statsig test projects")
	sdkKey, ok := os.LookupEnv("statsig_server_key")
	if !ok {
		panic("Expected statsig_server_key for logging test events")
	}
	statsig.Initialize(sdkKey)
	statsig.LogEvent(statsig.Event{EventName: "test_event_1", User: statsig.User{UserID: "test_user"}})
	statsig.Shutdown()
	println("Done")
}
