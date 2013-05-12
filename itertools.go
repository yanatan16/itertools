// Package itertools provides a translation of the python standard library module itertools.
// Many of the functions have been brought over, althought not all.
// In this implementation, chan interface{} has been used as all iterators; if more specific types are necessary,
// feel free to copy the code to your project to be implemented with more specific types.
package itertools

type Iter chan interface{}
type Predicate func (interface{}) bool
type Mapper func (interface{}) interface{}
type MultiMapper func (...interface{}) interface{}
type Reducer func (memo interface{}, element interface{}) interface{}

func New(els ... interface{}) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Int64(els ... int64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Int32(els ... int32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Float64(els ... float64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Float32(els ... float32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Uint(els ... uint) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}
func Uint64(els ... uint64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Uint32(els ... uint32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

// Count from i to infinity
func Count(i int) Iter {
	c := make(Iter)
	go func () {
		for ; true; i++ {
			c <- i
		}
	}()
	return c
}

// Cycle through an iterator infinitely (requires memory)
func Cycle(it Iter) Iter {
	c, a := make(Iter), make([]interface{}, 0, 1)
	go func () {
		for el := range it {
			a = append(a, el)
			c <- el
		}
		for {
			for _, el := range a {
				c <- el
			}
		}
	}()
	return c
}

// Repeat an element n times or infinitely
func Repeat(el interface{}, n ...int) Iter {
	c := make(Iter)
	go func () {
		for i := 0; len(n) == 0 || i < n[0]; i++ {
			c <- el
		}
		close(c)
	}()
	return c
}

// Chain together multiple iterators
func Chain(its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for _, it := range its {
			for el := range it {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Elements after pred(el) == true
func DropWhile(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if drop := pred(el); !drop {
				c <- el
				break
			}
		}
		for el := range it {
			c <- el
		}
		close(c)
	}()
	return c
}


// Elements before pred(el) == false
func TakeWhile(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if take := pred(el); take {
				c <- el
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Filter out any elements where pred(el) == false
func Filter(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if keep := pred(el); keep {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Filter out any elements where pred(el) == true
func FilterFalse(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if drop := pred(el); !drop {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Sub-iterator from start (inclusive) to [stop (exclusive) every [step (default 1)]]
func Slice(it Iter, startstopstep...int) Iter {
	start, stop, step := 0, 0, 1
	if len(startstopstep) == 1 {
		start = startstopstep[0]
	} else if len(startstopstep) == 2 {
		start, stop = startstopstep[0], startstopstep[1]
	} else	if len(startstopstep) >= 3 {
		start, stop, step = startstopstep[0], startstopstep[1], startstopstep[2]
	}

	c := make(Iter)
	go func () {
		i := 0
		// Start
		for el := range it {
			if i >= start {
				c <- el // inclusive
				break
			}
			i += 1
		}

		// Stop
		i, j := i + 1, 1
		for el := range it {
			if stop > 0 && i >= stop {
				break
			} else if j % step == 0 {
				c <- el
			}

			i, j = i + 1, j + 1
		}

		close(c)
	}()
	return c
}

// Map an iterator to fn(el) for el in it
func Map(fn Mapper, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			c <- fn(el)
		}
		close(c)
	}()
	return c
}

// Map p, q, ... to fn(pEl, qEl, ...)
// Breaks on first closed channel
func MultiMap(fn MultiMapper, its ...Iter) Iter {
	c := make(Iter)
	go func() {
Outer:
		for {
			els := make([]interface{}, len(its))
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					break Outer
				}
			}
			c <- fn(els...)
		}
		close(c)
	}()
	return c
}

// Map p, q, ... to fn(pEl, qEl, ...)
// Breaks on last closed channel
func MultiMapLongest(fn MultiMapper, its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for {
			els := make([]interface{}, len(its))
			n := 0
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					n += 1
				}
			}
			if n < len(its) {
				c <- fn(els...)
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Map an iterator if arrays to a fn(els...)
// Iter must be an iterator of []interface{} (possibly created by Zip)
// If not, Starmap will act like MultiMap with a single iterator
func Starmap(fn MultiMapper, it Iter) Iter {
	c := make(Iter)
	go func() {
		for els := range it {
			if elements, ok := els.([]interface{}); ok {
				c <- fn(elements...)
			} else {
				c <- fn(els)
			}
		}
		close(c)
	}()
	return c
}

// Zip up multiple interators into one
// Close on shortest iterator
func Zip(its ...Iter) Iter {
	c := make(Iter)
	go func() {
Outer:
		for {
			els := make([]interface{}, len(its))
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					break Outer
				}
			}
			c <- els
		}
		close(c)
	}()
	return c
}

// Zip up multiple iterators into one
// Close on longest iterator
func ZipLongest(its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for {
			els := make([]interface{}, len(its))
			n := 0
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					n += 1
				}
			}
			if n < len(its) {
				c <- els
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Reduce the iterator (aka fold) from the left
func Reduce(it Iter, red Reducer, memo interface{}) interface{} {
	for el := range it {
		memo = red(memo, el)
	}
	return memo
}