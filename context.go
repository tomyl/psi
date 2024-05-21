package psi

import (
	"context"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

type Context struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
	group  *errgroup.Group
}

func NewContext(parent context.Context) *Context {
	parent, cancel := context.WithCancelCause(parent)
	group, ctx := errgroup.WithContext(parent)
	group.SetLimit(runtime.GOMAXPROCS(-1))
	return &Context{
		ctx:    ctx,
		cancel: cancel,
		group:  group,
	}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *Context) SetLimit(n int) {
	c.group.SetLimit(n)
}

func (c *Context) Go(f func() error) {
	c.group.Go(f)
}

func (c *Context) Wait() error {
	return c.group.Wait()
}

func (c *Context) Cancel(err error) {
	c.cancel(err)
}
