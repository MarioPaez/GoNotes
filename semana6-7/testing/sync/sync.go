package sync

import "sync"

type Counter struct {
	mutex sync.Mutex
	Val   int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Val++
}

func (c *Counter) Value() int {
	return c.Val
}
