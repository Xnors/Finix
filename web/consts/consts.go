package consts

import "runtime"

const (
	PORT = "8080"
	// CLI_PATH = `..\cli\finix.exe`
)

func CLI_PATH() string {
	if runtime.GOOS == "linux" {
		return `./bin/finix`
	} else if runtime.GOOS == "windows" {
		return `.\bin\finix.exe`
	}
	return `../bin/finix`
}
