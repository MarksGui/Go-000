package rolling

type Bucket struct {
	Count  int64     // 桶累加值
	Points []float64 // 桶每次放入的数量
	next   *Bucket
}

//
func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

// 往偏移的桶里加入值
func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] = val
	b.Count++
}

// 重置桶
func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

// 返回指向的下一个桶
func (b *Bucket) Next() *Bucket {
	return b.next
}

// 窗口
type Window struct {
	window []Bucket
	size   int
}

type WindowOpts struct {
	Size int
}

// 构造方法
func NewWindow(opts WindowOpts) *Window {
	buckets := make([]Bucket, opts.Size)
	for offset := range buckets {
		buckets[offset].Points = make([]float64, 0)
		nextOffset := offset + 1
		if nextOffset == opts.Size {
			nextOffset = 0
		}
		// 创建循环链表
		buckets[offset].next = &buckets[nextOffset]
	}
	return &Window{window: buckets, size: opts.Size}
}

// 重置 bucket
func (w *Window) ResetBucket(offset int) {
	w.window[offset].Reset()
}

// 重置整个窗口
func (w *Window) ResetWindow() {
	for offset := range w.window {
		w.ResetBucket(offset)
	}
}

// 向窗口偏移的桶内加入数据
func (w *Window) Append(offset int, val float64) {
	w.window[offset].Append(val)
}

func (w *Window) Add(offset int, val float64) {
	// window 被重置后重新 append 元素
	if w.window[offset].Count == 0 {
		w.window[offset].Append(val)
		return
	}
	w.window[offset].Add(0, val)
}

// 返回窗口偏移的桶
func (w *Window) Bucket(offset int) Bucket {
	return w.window[offset]
}

// 返回窗口的长度
func (w *Window) Size() int {
	return w.size
}

func (w *Window) Iterator(offset int, count int) Iterator {
	return Iterator{
		count: count,
		cur:   &w.window[offset],
	}
}
