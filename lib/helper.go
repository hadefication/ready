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
