package limiter

import (
	"github.com/go-pay/limiter/group"
	"github.com/go-pay/limiter/rate"
)

var (
	defaultConfig = &Config{
		Rate:       1000,
		BucketSize: 1000,
	}
)

type Config struct {
	// per second request，0 不限流
	Rate int `json:"rate" yaml:"rate" toml:"rate"`

	// max size，桶内最大量
	BucketSize int `json:"bucket_size" yaml:"bucket_size" toml:"bucket_size"`
}

// 速率限制器
type RateLimiter struct {
	C            *Config
	LimiterGroup *group.RateGroup
}

func NewLimiter(c *Config) (rl *RateLimiter) {
	if c == nil {
		c = defaultConfig
	}
	rl = &RateLimiter{
		C: c,
		LimiterGroup: group.NewRateGroup(func() *rate.Limiter {
			return newLimiter(c)
		}),
	}
	return rl
}

func newLimiter(c *Config) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(c.Rate), c.BucketSize)
}
