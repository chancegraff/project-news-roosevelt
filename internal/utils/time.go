package utils

import "time"

// Tomorrow will return a time.Time instance for tomorrow
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}
