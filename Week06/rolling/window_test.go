package rolling

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t1 := time.Now()
	time.Sleep(1 * time.Microsecond)
	fmt.Println(int(time.Since(t1) / (100 * time.Millisecond)))
}
