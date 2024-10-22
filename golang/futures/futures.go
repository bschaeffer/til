package futures

import "sync"

func Do[T any](work func() (T, error)) Future[T] {
	f := &future[T]{}
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.res, f.err = work()
	}()
	return f
}

type Future[T any] interface {
	Await() (T, error)
}

type future[T any] struct {
	wg  sync.WaitGroup
	res T
	err error
}

func (f *future[T]) Await() (T, error) {
	f.wg.Wait()
	return f.res, f.err
}
