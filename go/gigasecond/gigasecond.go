package gigasecond

import (
	"fmt"
	"math"
	"time"
)

var (
	TestVersion           = 1
	Birthday    time.Time = time.Date(1977, time.June, 13, 0, 0, 0, 0, time.UTC)
	print                 = fmt.Println
)

func AddGigasecond(tt time.Time) time.Time {
	gigaseconds := time.Duration(math.Pow10(9)) * time.Second
	tt = tt.Add(gigaseconds)
	return tt
}
