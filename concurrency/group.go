package concurrency

type Group struct {
	sem chan struct{}
}

func NewGroup(max int) *Group {
	return &Group{sem: make(chan struct{}, max)}
}

func (g *Group) Go(f func()) {
	g.sem <- struct{}{} // wait until the number of goroutines is below the limit
	go func() {
		defer func() { <-g.sem }() // let another goroutine run
		f()
	}()
}

func (g *Group) Wait() {
	for i := 0; i < cap(g.sem); i++ {
		g.sem <- struct{}{}
	}
}
