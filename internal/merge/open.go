package merge

import (
	"os/exec"
	"runtime"
)

func OpenFolder(folderPath string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("explorer", folderPath).Start()
	case "darwin":
		exec.Command("open", folderPath).Start()
	default:
		exec.Command("xdg-open", folderPath).Start()
	}
}
