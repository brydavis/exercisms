package gigasecond

// package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// const (
// 	fmtD  = "2006-01-02"
// 	fmtDT = "2006-01-02T15:04:05"
// )

// var (
// 	TestVersion           = 1
// 	Birthday    time.Time = time.Date(1977, time.June, 13, 0, 0, 0, 0, time.UTC)
// 	print                 = fmt.Println
// )

// func AddGigasecond(tt time.Time) {

// 	gigaseconds := time.Duration(math.Pow10(9)) * time.Second

// 	print("Before: ", tt.Format(fmtD))
// 	tt = tt.Add(gigaseconds)
// 	print("After: ", tt.Format(fmtDT))

// }

// func main() {
// 	AddGigasecond(Birthday)
// }
