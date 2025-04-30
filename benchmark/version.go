package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// getLatestVersion fetches the latest version tag for a given Go module URL.
func getLatestVersion(moduleURL string) (string, error) {
	// Convert https://github.com/user/repo to github.com/user/repo
	modulePath := strings.TrimPrefix(moduleURL, "https://")
	cmd := exec.Command("go", "list", "-m", "-json", modulePath+"@latest")
	output, err := cmd.Output()
	if err != nil {
		// Combine stderr with error message if available
		if exitError, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("go list command failed for %s: %v\nStderr: %s", modulePath, err, string(exitError.Stderr))
		}
		return "", fmt.Errorf("go list command failed for %s: %v", modulePath, err)
	}

	var modInfo struct {
		Version string
	}
	if err := json.Unmarshal(output, &modInfo); err != nil {
		return "", fmt.Errorf("failed to parse go list output for %s: %v\nOutput: %s", modulePath, err, string(output))
	}

	if modInfo.Version == "" {
		return "", fmt.Errorf("no version found for %s", modulePath)
	}

	return modInfo.Version, nil
}
