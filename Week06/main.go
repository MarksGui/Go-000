package main

import (
	"Go-000/Week06/rolling"
	"fmt"
	"time"
)

func main() {
	rollingCounter := rolling.NewRollingCounter(
		rolling.NewWindow(rolling.WindowOpts{Size: 10}),
		rolling.RollingCounterOpts{
			BucketDuration: 100 * time.Millisecond,
		},
	)

	rollingCounter.Add(4)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
	fmt.Println("----------------")
	rollingCounter.Add(6)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
	fmt.Println("----------------")
	time.Sleep(200 * time.Millisecond)
	rollingCounter.Add(1)
	rollingCounter.Add(3)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
	fmt.Println("----------------")
	time.Sleep(100 * time.Millisecond)
	rollingCounter.Add(10)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
	fmt.Println("----------------")
	time.Sleep(500 * time.Millisecond)
	rollingCounter.Add(100)
	fmt.Printf("rolling counter avg:%f,Value:%f\n", rollingCounter.Reduce(rolling.Avg), rollingCounter.Reduce(rolling.Sum))
}
