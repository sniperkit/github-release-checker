package readme

import (
	"sync"
)

type Chan <-chan *Readme
type chanW chan<- *Readme
type chanRW chan *Readme

// TODO(leon): This is ugly
func onlyReadable(in chanRW) <-chan *Readme {
	return in
}

// TODO(leon): This is ugly
func onlyWritable(in chanRW) chan<- *Readme {
	return in
}

func CloneChan(in Chan, l int) []Chan {
	// Initialize clones
	outRW := make([]chanRW, l)
	for i := range outRW {
		outRW[i] = make(chanRW)
	}

	go func() {
		// Populate clones
		wgs := make([]sync.WaitGroup, l)
		for t := range in {
			for i, c := range outRW {
				wg := &wgs[i]
				wg.Add(1)
				go func(t *Readme, c chanRW) {
					defer wg.Done()
					c <- t
				}(t, c)
			}
		}
		// Close clones
		for i := range outRW {
			go func(i int) {
				wgs[i].Wait()
				close(outRW[i])
			}(i)
		}
	}()

	out := make([]Chan, l)
	for i := range outRW {
		out[i] = onlyReadable(outRW[i])
	}
	return out
}
