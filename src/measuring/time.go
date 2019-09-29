package measuring

import (
	"log"
	"time"
)

func ElapsedTime(start time.Time, name string) time.Duration {
	result := time.Since(start)
	log.Println(name, result)
	return result
}
