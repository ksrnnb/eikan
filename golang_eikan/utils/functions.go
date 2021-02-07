package utils

import "time"

// CurrentTime returns current time based on RFC3339
func CurrentTime() time.Time {
	return time.Now()
}
