package rateLimiter

import (
	"sync"
	"time"
)

type Bucket struct {
	IP        string
	CreatedAt time.Time
	LastReset time.Time
	Capacity  int64
	Available int64
	mutex     sync.Mutex
}

func NewBucket(capacity int64, ip string) *Bucket {
	return &Bucket{
		IP:        ip,
		CreatedAt: time.Now(),
		LastReset: time.Now(),
		Capacity:  capacity,
		Available: capacity,
	}
}

func (bucket *Bucket) refill() {
	bucket.Available = bucket.Capacity
	bucket.LastReset = time.Now()
}

func (bucket *Bucket) IsAllowed(tokens int64) (bool, time.Duration, int64) {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock()

	var tryAfter time.Duration

	elapsed := time.Since(bucket.LastReset)
	if elapsed >= 60*time.Second {
		bucket.refill()
	} else {
		tryAfter = (60 * time.Second) - elapsed
		// tryAfter = (60 * time.Second) - time.Duration(elapsed.Seconds())
	}

	if bucket.Available >= tokens {
		bucket.Available = bucket.Available - tokens
		return true, 0, bucket.Available
	}
	return false, tryAfter, 0
}
