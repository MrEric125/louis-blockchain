package cache

type Iterator struct {
	Val interface{}

	// expire 过期时间
	Expire int64
}
