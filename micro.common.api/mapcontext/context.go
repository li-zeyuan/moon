package mapcontext

import (
	"context"
	"sync"
	"time"
)

type mapCtx struct {
	Ctx context.Context
	Maps sync.Map
}

func NewContext(ctx context.Context) *mapCtx {
	return &mapCtx{
		Ctx: ctx,
		Maps: sync.Map{},
		}
}

func (*mapCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*mapCtx) Done() <-chan struct{} {
	return nil
}

func (*mapCtx) Err() error {
	return nil
}

func (c *mapCtx) Value(key interface{}) interface{} {
	str, ok := key.(string)
	if !ok {
		return nil
	}
	value, _ := c.Maps.Load(str)
	return value
}

func (c *mapCtx) Set(key ,value interface{}) {
	c.Maps.Store(key, value)
}

func (c *mapCtx) Del(key interface{}) {
	c.Maps.Delete(key)
}