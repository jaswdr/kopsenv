package internal

import (
	"time"
)

type Release struct {
	Tag string `json:"tag"`
}

type Data struct {
	Timestamp time.Time `json:"timestamp"`
	Releases  []Release `json:"releases"`
}
