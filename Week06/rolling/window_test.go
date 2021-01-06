package rolling

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(int(time.Since(t1) / (3 * time.Millisecond)))
}
