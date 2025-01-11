package core

import (
	"runtime"
	"sync"

	"github.com/buke/quickjs-go"
)

type Task func(*quickjs.Context) any

type TaskWrapper struct {
	task     Task
	receiver chan any
}

type QuickJSWrapper struct {
	wg    sync.WaitGroup
	rt    quickjs.Runtime
	ctx   *quickjs.Context
	queue chan TaskWrapper
}

func (q *QuickJSWrapper) Done() {
	q.wg.Done()
}

func (q *QuickJSWrapper) Add(i int) {
	q.wg.Add(i)
}

func (q *QuickJSWrapper) Wait() {
	q.wg.Wait()
	close(q.queue)
}

func NewQuickJSWrapper() QuickJSWrapper {
	rt := quickjs.NewRuntime()
	ctx := rt.NewContext()

	return QuickJSWrapper{
		rt:    rt,
		ctx:   ctx,
		queue: make(chan TaskWrapper, 50),
	}
}

func (q *QuickJSWrapper) Start() {
	go func() {
		// Lock OS thread so calls to contexts are all from the same thread
		// QuickJS is not thread safe with respect to runtimes, so ensure that all
		// code in this goroutine runs
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		defer q.Done()
		for w := range q.queue {
			w.receiver <- w.task(q.ctx)
			// Ensure that Event bubbles are processed.
			q.ctx.Loop()
			q.Done()
		}
	}()
}

// Used for thread safe runtime for calling within QuickJS context.
// Only use context in here. Just note that you MUST
// use another `With` to obtain context is work is done elsewhere for AsyncFunctions created
// to call promise again afterwords. Also note that AsyncFunction and Function bodies are blocking,
// unless using a GoRoutine in an AsyncFunction and latter calling the
// Be sure to `Add`/`Done` for the loop so that the event loop knows there's work to be waited on.
func (q *QuickJSWrapper) With(task Task) <-chan any {
	q.Add(1)
	receiver := make(chan any, 1)

	q.queue <- TaskWrapper{
		task:     task,
		receiver: receiver,
	}

	return receiver
}
