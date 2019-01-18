package osascript

import (
	"fmt"
	"os/exec"
)

const scriptCmd = "tell front window to make new tab at after (get active tab) with properties {URL:\"%s\"}"

func OpenNewChromeTab(url string) error {
	cmd := exec.Command(
		"osascript",
		"-e", "tell application \"Google Chrome\"",
		"-e", "activate",
		"-e", fmt.Sprintf(scriptCmd, url),
		"-e", "end tell")

	return cmd.Run()
}
