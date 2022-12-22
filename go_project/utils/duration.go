package utils

import (
	"strings"
	"time"
)

func parseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dura, error := time.ParseDuration(d)
	
}
