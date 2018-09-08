package runenv

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

// EnsureSupportedRunEnv checks to make sure that the program is running on a supported operating system and
// linux distribution.
//
// An error is returned if details about the current OS cannot be retrieved or if the current OS is not supported.
func EnsureSupportedRunEnv() error {
	// Check running linux
	if runtime.GOOS != "linux" {
		return fmt.Errorf("tool can only run on linux, currently running on \"%s\"", runtime.GOOS)
	}

	// Check running arch linux
	dist, err := GetDistributionDesc()
	if err != nil {
		return fmt.Errorf("failed to get the name of the current linux distribution: %s", err.Error())
	}

	if dist != "Arch Linux" {
		return fmt.Errorf("tool can only run on the Arch Linux distribution, current linux distribution: %s", dist)
	}

	return nil
}

// LSBDistributionDescRegexp is the regular expression used to extract the distribution id field from an lsb_release
// execution output
var LSBDistributionDesxRegexp *regexp.Regexp = regexp.MustCompile("Description:\\s+([A-Za-z ]+)")

// GetDistributionDesc retrieves the name of the current linux distribution using the lsb_release tool.
//
// An error is returned if the lsb_release tool fails to run.
func GetDistributionDesc() (string, error) {
	cmd := exec.Command("lsb_release", "-a")

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running lsb_release tool: %s", err.Error())
	}

	outStr := string(out)

	outMatches := LSBDistributionDesxRegexp.FindStringSubmatch(outStr)
	if outMatches == nil {
		return "", fmt.Errorf("lsb_release output not in expected format, was: %s", outStr)
	}

	return outMatches[1], nil
}
