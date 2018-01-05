package timeCost

import "time"

func Cost(start time.Time)int{
	return int(time.Since(start))
}
