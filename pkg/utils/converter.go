package utils

import "time"

func RemainingTimeForExpiration(epoch int64) time.Duration {
	currentTime := time.Now().Unix()
	if epoch < currentTime {
		return 0
	}
	return time.Duration(epoch-currentTime) * time.Second
}
