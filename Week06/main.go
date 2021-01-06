package main

import (
	"Go-000/Week06/rolling"
	"fmt"
	"time"
)

func main() {
	w := rolling.NewWindow(rolling.WindowOpts{Size: 10})
	opt := rolling.RollingCounterOpts{BucketDuration: 100 * time.Millisecond}
	rollingCounter := rolling.NewRollingCounter(w, opt)

	rollingCounter.Add(4)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
	rollingCounter.Add(6)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))

	time.Sleep(200 * time.Millisecond)
	rollingCounter.Add(1)
	rollingCounter.Add(3)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))

	time.Sleep(100 * time.Millisecond)
	rollingCounter.Add(10)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))

	time.Sleep(500 * time.Millisecond)
	rollingCounter.Add(100)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
}
