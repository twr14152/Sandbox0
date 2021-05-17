//https://play.golang.org/p/nJnS5ZU1Ydi

package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Printf("DayTime: %v\n", dt)
	fmt.Printf("DayTime Type: %T\n", dt)
	fmt.Printf("Today is: %v\n", dt.Weekday())
	
	y,m,d := dt.Date()
	fmt.Printf("Date: %v %v %v \n", m,d,y)
	
	yr, wk := dt.ISOWeek()
	fmt.Printf("Week No.: %v in %v\n", wk, yr)
	
	dy := dt.YearDay()
	fmt.Printf("Day No.: %v\n", dy)
}
/*
DayTime: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
DayTime Type: time.Time
Today is: Tuesday
Date: November 10 2009 
Week No.: 46 in 2009
Day No.: 314
*/
