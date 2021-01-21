package lib

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/otiai10/copy"
)

// BackupDockerComposeFile backups the docker-compose.yml
func BackupDockerComposeFile(path string, runtimePath string) {
	dockerComposeFile := path + "/docker-compose.yml"

	if _, err := os.Stat(dockerComposeFile); !os.IsNotExist(err) {
		t := time.Now()
		formatted := fmt.Sprintf("%d%02d%02dT%02d%02d%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		os.Rename(dockerComposeFile, dockerComposeFile+"."+formatted+".bak")
	}

	copy.Copy(runtimePath+"/docker-compose.yml", dockerComposeFile)
}

// IsRuntimeURL test if supplied URL is an actual URL
func IsRuntimeURL(urlToTest string) bool {
	_, err := url.ParseRequestURI(urlToTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(urlToTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

//
var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red
)

//
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

// Color is
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
