package utils

import "time"

// CurrentTime returns current time based on RFC3339
func CurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
