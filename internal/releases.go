package internal

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	RAW_TAGS_LIST_URL string = "https://raw.githubusercontent.com/jaswdr/kopsenv/refs/heads/master/data/tags"
	DOWNLOAD_URL      string = "https://github.com/kubernetes/kops/releases/download/%s/kops-%s-%s"
)

func buildReleaseFromTag(tag string) Release {
	versionSplit := strings.Split(tag, ".")

	major := strings.TrimLeft(versionSplit[0], "v")
	log := logrus.WithField("majorVersion", major)
	majorAsInt, err := strconv.Atoi(major)
	if err != nil {
		log.Fatal("Major version is not a valid integer")
	}

	minor := versionSplit[1]
	log = log.WithField("minorVersion", minor)

	patch := strings.Join(versionSplit[2:], ".")
	log = log.WithField("patch", patch)

	minorAsInt, err := strconv.Atoi(minor)
	if err != nil {
		log.Fatal("Minor version is not a valid integer")
	}

	isAlpha := strings.Contains(tag, "alpha")
	isBeta := strings.Contains(tag, "beta")
	patchRelease := 0
	if isAlpha || isBeta {
		patchSplit := strings.Split(patch, ".")
		if len(patchSplit) > 1 {
			// beta.X format
			patchRelease, err = strconv.Atoi(patchSplit[1])
			if err != nil {
				log.WithError(err).Fatal("Invalid patch release format")
			}
		} else {
			// betaX format
			tempSplit := strings.Split(patch, "-")
			temp := strings.TrimLeft(tempSplit[1], "alpha")
			temp = strings.TrimLeft(temp, "beta")
			patchRelease, err = strconv.Atoi(temp)
			if err != nil {
				log.WithError(err).Fatal("Invalid patch release format")
			}
		}
	}

	return Release{
		Tag:          tag,
		Major:        majorAsInt,
		Minor:        minorAsInt,
		Patch:        patch,
		IsAlpha:      isAlpha,
		IsBeta:       isBeta,
		PatchRelease: patchRelease,
	}
}

func applySort(releases []Release) {
	sort.Slice(releases, func(i, j int) bool {
		left := releases[i]
		right := releases[j]

		if left.Major == right.Major {
			if left.Minor == right.Minor {
				if (left.IsAlpha && right.IsAlpha) || (left.IsBeta && right.IsBeta) {
					return left.PatchRelease > right.PatchRelease
				}

				if left.IsBeta && right.IsAlpha {
					return true
				} else {
					return false
				}
			}
			return left.Minor > right.Minor
		}

		return left.Major > right.Major
	})
}

func GetReleases() (result []Release) {
	resp, err := http.Get(RAW_TAGS_LIST_URL)
	if err != nil {
		log.WithField("url", RAW_TAGS_LIST_URL).WithError(err).Fatal("Failed to read list of tags from remote file")
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to response body from list of tags response")
	}

	tags := strings.Split(strings.TrimSpace(string(content)), "\n")
	result = make([]Release, len(tags))
	for i, tag := range tags {
		result[i] = buildReleaseFromTag(tag)
	}

	applySort(result)
	return result
}

func Download(version string) {
	runtimeOs := runtime.GOOS
	runtimeArchitecture := runtime.GOARCH
	url := fmt.Sprintf(DOWNLOAD_URL, version, runtimeOs, runtimeArchitecture)
	log := logrus.WithFields(logrus.Fields{
		"version": version,
		"os":      runtimeOs,
		"arch":    runtimeArchitecture,
		"url":     url,
	})

	log.Info("Downloading version from GitHub")

	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Fatal("Failed to download version")
	}

	log.Info("Response from remote server received")
	log.Info("Reading response from remote server")
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to read response body")
	}

	log = log.WithField("responseSize", len(content))
	log.Info("Entire response was successfully read")

	release := buildReleaseFromTag(version)
	SaveRelease(release, content)
}
