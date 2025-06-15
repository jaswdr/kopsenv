package internal

import (
	"time"
)

type Release struct {
	Tag          string `json:"tag"`
	Major        int    `json:"major"`
	Minor        int    `json:"minor"`
	Patch        string `json:"patch"`
	IsAlpha      bool   `json:"is_alpha"`
	IsBeta       bool   `json:"is_beta"`
	PatchRelease int    `json:"patch_release"`
}

func (r Release) String() string {
	return r.Tag
}

type Data struct {
	Timestamp time.Time `json:"timestamp"`
	Releases  []Release `json:"releases"`
}
