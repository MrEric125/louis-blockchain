package cache

import (
	"context"
	"time"
)

type sentinel struct {

	// 间隔时间 0 不开启哨兵
	interval time.Duration

	// ctx
	ctx context.Context
	// 哨兵周期执行的函数
	fn func()
}

func (s *sentinel) Start() {
	if s.interval <= 0 {
		return
	}
	tick := time.NewTicker(s.interval)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			s.fn()
		case <-s.ctx.Done():
			return
		}
	}

}
