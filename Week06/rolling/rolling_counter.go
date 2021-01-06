package rolling

import (
	"sync"
	"time"
)

type RollingPolicy struct {
	mu     sync.RWMutex // 读写锁
	size   int          // 窗口长度
	window *Window      // 窗口
	offset int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

// 滑动窗口配置
type RollingCounterOpts struct {
	BucketDuration time.Duration
}

func NewRollingCounter(window *Window, opts RollingCounterOpts) *RollingPolicy {
	return &RollingPolicy{
		window: window,
		size:   window.Size(),
		offset: 0,

		bucketDuration: opts.BucketDuration,
		lastAppendTime: time.Now(),
	}
}

// 计算当前时间到最后一次追加时间之间的桶间隔
func (r *RollingPolicy) timespan() int {
	v := int(time.Since(r.lastAppendTime) / r.bucketDuration)
	if v > -1 {
		return v
	}
	return r.size
}

// 向华东窗口里面添加元素
func (r *RollingPolicy) add(f func(offset int, val float64), val float64) {
	r.mu.Lock()
	timespan := r.timespan()
	if timespan > 0 {
		r.lastAppendTime = r.lastAppendTime.Add(time.Duration(timespan * int(r.bucketDuration)))
		offset := r.offset
		// reset the expired buckets
		s := offset + 1
		if timespan > r.size {
			timespan = r.size
		}
		e, e1 := s+timespan, 0 // e: reset offset must start from offset+1
		if e > r.size {
			e1 = e - r.size
			e = r.size
		}
		for i := s; i < e; i++ {
			r.window.ResetBucket(i)
			offset = i
		}
		for i := 0; i < e1; i++ {
			r.window.ResetBucket(i)
			offset = i
		}
		r.offset = offset
	}
	f(r.offset, val)
	r.mu.Unlock()
}

// Append appends the given points to the window.
func (r *RollingPolicy) Append(val float64) {
	r.add(r.window.Append, val)
}

// Add adds the given value to the latest point within bucket.
func (r *RollingPolicy) Add(val float64) {
	r.add(r.window.Add, val)
}

// Reduce applies the reduction function to all buckets within the window.
func (r *RollingPolicy) Reduce(f func(Iterator) float64) (val float64) {
	r.mu.RLock()
	timespan := r.timespan()
	if count := r.size - timespan; count > 0 {
		offset := r.offset + timespan + 1
		if offset >= r.size {
			offset = offset - r.size
		}
		val = f(r.window.Iterator(offset, count))
	}
	r.mu.RUnlock()
	return val
}
