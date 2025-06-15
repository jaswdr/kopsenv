package internal

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func cacheBaseDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("Error when building base cache dir")
	}

	return strings.Join([]string{homeDir, ".kopsenv"}, string(os.PathSeparator))
}

func cacheReleasesFile() string {
	cacheDir := cacheBaseDir()
	return strings.Join([]string{cacheDir, "releases.json"}, string(os.PathSeparator))
}

func saveData(data Data) {
}

func loadData() *Data {
	return nil
}
