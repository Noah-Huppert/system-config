package main

import (
	"github.com/Noah-Huppert/system-config/runenv"

	"github.com/Noah-Huppert/golog"
)

func main() {
	// Setup logger
	logger := golog.NewStdLogger("syscfg")

	// Check if running on a supported OS
	err := runenv.EnsureSupportedRunEnv()
	if err != nil {
		logger.Fatalf("failed to verify the tool is running on a supported operating system: %s", err.Error())
	}
}
