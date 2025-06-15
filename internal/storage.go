package internal

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR           = ".kopsenv"
	RELEASES_FOLDER_NAME = "releases"
	BIN_FOLDER_NAME      = "bin"
	BIN_FILE_NAME        = "kops"
)

func configDirPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Fatal("Failed to get home directory")
	}

	return strings.Join([]string{homeDir, CONFIG_DIR}, string(os.PathSeparator))
}

func versionPathForRelease(release Release) string {
	return strings.Join([]string{
		configDirPath(), RELEASES_FOLDER_NAME, release.Tag,
	}, string(os.PathSeparator))
}

func SaveRelease(release Release, content []byte) {
	destPath := versionPathForRelease(release)
	log := logrus.WithFields(logrus.Fields{
		"release":     release,
		"contentSize": len(content),
		"destination": destPath,
	})

	destBase := filepath.Dir(destPath)
	log = log.WithField("destinationBase", destBase)

	err := os.MkdirAll(destBase, 0700)
	if err != nil {
		log.WithError(err).Fatal("Failed to create releases directory")
	}

	destFile, err := os.OpenFile(destPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		log.WithError(err).Fatal("Failed to create destination file")
	}

	destFile.Write(content)
	destFile.Close()

	log = log.WithField("tempFileDeleted", true)
	log.Info("Version successfully saved locally")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func IsVersionAvailable(version string) bool {
	release := buildReleaseFromTag(version)
	versionPath := versionPathForRelease(release)
	return fileExists(versionPath)
}

func LinkVersion(version string) {
	release := buildReleaseFromTag(version)
	orig := versionPathForRelease(release)
	destBaseDir := strings.Join([]string{configDirPath(), BIN_FOLDER_NAME}, string(os.PathSeparator))
	dest := strings.Join([]string{destBaseDir, BIN_FILE_NAME}, string(os.PathSeparator))

	log := logrus.WithFields(logrus.Fields{
		"origin":          orig,
		"destinationBase": destBaseDir,
		"destination":     dest,
	})

	err := os.MkdirAll(destBaseDir, 0700)
	if err != nil {
		log.WithError(err).Fatal("Failed to create destination base folder")
	}

	if fileExists(dest) {
		err = os.Remove(dest)
		if err != nil {
			log.WithError(err).Fatal("Failed to remove destination file")
		}
	}

	err = os.Symlink(orig, dest)
	if err != nil {
		log.WithError(err).Fatal("Failed to create symlink")
	}
}
