package cache

import (
	"context"
	"fmt"
	"louis/options"
	"sync"
	"time"
)

type BlackCache struct {
	*cache
}

type cache struct {
	sync.RWMutex

	// sentinel 维护一个周期循环的任务
	*sentinel

	// defaultExpire 默认超时时间
	defaultExpire time.Duration

	// member 维护map存储kv关系
	member map[string]Iterator

	// capture 捕获删除对象时间 会返回kv值用于用户自定义处理
	capture func(k string, v interface{})

	cancel context.CancelFunc
}

func (c *cache) DeleteExpire() {

}

func NewCache(options ...options.Option) BlackCache {
	ctx, cancel := context.WithCancel(context.Background())

	c := &Config{
		defaultExpire: 0,
		interval:      0,
		capture: func(k string, v interface{}) {
			fmt.Printf("delete k:%s v:%v\n", k, v)
		},
	}
	for _, option := range options {
		option(c)
	}
	obj := &cache{
		defaultExpire: c.defaultExpire,
		capture:       c.capture,
		cancel:        cancel,
	}
	if c.member == nil {
		c.member = map[string]Iterator{}
	}
	if c.fn == nil {
		c.fn = obj.DeleteExpire
	}
	obj.member = c.member
	obj.sentinel = &sentinel{
		interval: c.interval,
		ctx:      ctx,
		fn:       c.fn,
	}
	go obj.sentinel.Start()
	return BlackCache{obj}
}
