package group

import (
	"sync"

	"github.com/go-pay/limiter/rate"
	"github.com/go-pay/smap"
)

// RateGroup 速率Group
type RateGroup struct {
	new func() *rate.Limiter
	rgs smap.Map[string, *rate.Limiter]
	sync.RWMutex
}

// NewRateGroup 新建RateGroup
func NewRateGroup(new func() *rate.Limiter) (rg *RateGroup) {
	if new == nil {
		panic("RateGroup: can't assign a nil to the new function")
	}
	return &RateGroup{new: new}
}

// Get 获取Limiter，如果没有则新建
func (r *RateGroup) Get(key string) *rate.Limiter {
	rg, ok := r.rgs.Load(key)
	if !ok {
		r.RLock()
		newRg := r.new
		r.RUnlock()
		rg = newRg()
		r.rgs.Store(key, rg)
	}
	return rg
}
