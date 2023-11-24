package limiter

import (
	"log"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	l := NewLimiter(&Config{
		Rate:       10,
		BucketSize: 10,
	})
	for i := 0; i < 200; i++ {
		time.Sleep(10 * time.Millisecond)
		lr := l.LimiterGroup.Get("/test/limiter")
		if !lr.Allow() {
			log.Printf("[%d] limited wait...", i)
			continue
		}
		log.Printf("[%d] passed", i)
	}
}
