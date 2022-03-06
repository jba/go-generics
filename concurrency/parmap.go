package concurrency

import (
	"context"
	"sync"
)

// type Transformer[A, B any] struct {
// 	sem chan struct{}
// }

// func NewTransformer[A, B any](max int) Transformer[A, B] {
// 	if max <= 0 {
// 		return &Transformer[A, B]{}
// 	}
// 	return &Transformer[A, B]{sem: make(chan struct{}, max)}
// }

// func (t *Transformer[A, B]) Slice(in []A, f func(A) (B, error)) ([]B, error) {

func TransformSlice[A, B any](ctx context.Context, in []A, max int, f func(context.Context, A) (B, error)) ([]B, error) {
	type el struct {
		i int
		a A
	}

	if max < 0 {
		max = len(in)
	}
	elc := make(chan el)

	errc := make(chan error, 1)
	senderr := func(err error) {
		select {
		case errc <- err:
		default:
		}
	}

	out := make([]B, len(in))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for e := range elc {
				select {
				case <-ctx.Done():
					senderr(ctx.Err())
					return
				default:
					b, err := f(ctx, e.a)
					if err != nil {
						cancel()
						senderr(err)
						return
					}
					out[e.i] = b
				}
			}
		}()
	}

	for i, a := range in {
		elc <- el{i, a}
	}
	close(elc)

	wg.Wait()
	select {
	case err := <-errc:
		return nil, err
	default:
		return out, nil
	}
}
