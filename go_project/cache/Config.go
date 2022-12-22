package cache

import (
	"louis/options"
	"time"
)

type Config struct {
	// defaultExpire 默认超时时间
	defaultExpire time.Duration

	// interval 间隔时间
	interval time.Duration
	// fn 哨兵周期执行的函数
	fn func()

	// capture 捕获删除对象时间 会返回kv值用于用户自定义处理
	capture func(k string, v interface{})

	member map[string]Iterator
}

// SetDefaultExpire 设置默认的超时时间
func SetDefaultExpire(expire time.Duration) options.Option {
	return func(c interface{}) {
		c.(*Config).defaultExpire = expire
	}
}
