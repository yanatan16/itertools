# Itertools for golang

This package is a translation of the python `itertools` module. It includes all the usual suspects except Tee and including Reduce. All iterators are `chan interface{}` which allows some type ambiguity for these generic functions. It would be completely ok, however, to reproduce these functions in your package for your type-specific iterators such as `chan MyStruct`. I did this mostly as a thought exercise on converting python generators to Go.

Full documentation is available on [godoc](http://godoc.org/github.com/yanatan16/itertools).

# Implemented Functions

- Infinite Iterator Creators
		- `Count(i)` - Infinite count from i
		- `Cycle(it)` - Infinite cycling of it (requires memory)
		- `Repeat(element [, n])` - Repeat element n times (or infinitely)
- Finite Iterator Creators
		- `New(elements ...)` - Create from `interface{}` elements
		- `Int32(elements ...)` - Create from `int32` elements
		- `Int64(elements ...)` - Create from `int64` elements
		- `Uint(elements ...)` - Create from `uint` elements
		- `Uint32(elements ...)` - Create from `uint32` elements
		- `Uint64(elements ...)` - Create from `uint64` elements
		- `Float32(elements ...)` - Create from `float32` elements
		- `Float64(elements ...)` - Create from `float64` elements
- Iterator Modifiers
		- `Chain(iters...)` - Chain together multiple iterators
		- `DropWhile(predicate, iter)` - Drop elements until `predicate(el) == false`
		- `TakeWhile(predicate, iter)` - Take elements until `predicate(el) == false`
		- `Filter(predicate, iter)` - Filter out elements when `predicate(el) == false`
		- `FilterFalse(predicate, iter)` - Filter out elements when `predicate(el) == true`
		- `Slice(iter, start[, stop[, step]])` - Drop elements until the start (0-based index). Stop upon stop (exclusive) unless not given. Step is 1 unless given.
		- `Map(mapper, iter)` - Map the iterator
		- `MultiMap(multiMapper, iters...)` - Map all the iterators as `multiMap(elements...)`. Stop on shortest iterator.
		- `MultiMapLongest(multiMapper, iters...)` - Map all the iterators as `multiMap(elements...)`. Stop on longest iterator. Shorter iterators are filled with `nil` after they are exhausted.
		- `Starmap(multiMapper, iter)` - If iter is an iterator of `[]interface{}`, then expand it into the `multiMapper`.
		- `Zip(iters...)` - Zip multiple iterators together
		- `ZipLongest(iters...)` - Zip multiple iterators together. Take the longest. Shorter ones are appended with `nil`.
		- `Reduce(iter, reducer, memo)` - Reduce (or Foldl) across the iterator.

# License

Copyright (c) 2013 Jon Eisen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.