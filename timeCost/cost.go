package timeCost

import "time"

func Cost(start time.Time)time.Duration{
	return time.Since(start)
}
