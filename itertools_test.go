package itertools

import (
	"testing"
	"reflect"
)

// Test iterators for element equality. Allow it1 to be longer than it2
func testIter(t *testing.T, it1, it2 Iter) {
	t.Log("Start")
	for el1 := range it1 {
		if el2, ok := <- it2; !ok {
			t.Error("it2 shorter than it1!", el1)
			return
		} else if !reflect.DeepEqual(el1, el2) {
			t.Error("Elements are not equal", el1, el2)
		} else {
			t.Log(el1, el2)
		}
	}
	t.Log("Stop")
}

// Test iterators for element equality. Don't allow it1 to be longer than it2
func testIterEq(t *testing.T, it1, it2 Iter) {
	t.Log("Start")
	for el1 := range it1 {
		if el2, ok := <- it2; !ok {
			t.Error("it2 shorter than it1!", el1)
			return
		} else if !reflect.DeepEqual(el1, el2) {
			t.Error("Elements are not equal", el1, el2)
		} else {
			t.Log(el1, el2)
		}
	}
	if el2, ok := <- it2; ok {
		t.Error("it1 shorter than it2!", el2)
	}
	t.Log("Stop")
}

func TestCount(t *testing.T) {
	testIter(t, Int(1,2,3,4,5,6,7,8,9), Count(1))
}

func TestCycle(t *testing.T) {
	testIter(t, String("a", "b", "ccc", "a", "b", "ccc", "a"), Cycle(String("a", "b", "ccc")))
}

func TestRepeat(t *testing.T) {
	testIterEq(t, Uint64(100, 100, 100, 100), Repeat(uint64(100), 4))
	testIter(t, Uint64(100, 100, 100, 100), Repeat(uint64(100)))
}

func TestChain(t *testing.T) {
	testIterEq(t, Int32(1,2,3,4,5,5,4,3,2,1,100), Chain(Int32(1,2,3,4,5), Int32(5,4,3,2,1), Int32(100)))
}

func TestCompress(t *testing.T) {
	testIter(t, Int(1,3,4,6), Compress(Count(0), Bool(false, true, false, true, true, false, true)))
}


func TestDropWhile(t *testing.T) {
	pred := func (i interface{}) bool {
		return i.(int) < 10
	}
	testIter(t, Int(10,11,12,13,14,15), DropWhile(pred, Count(0)))
}

func TestTakeWhile(t *testing.T) {
	pred := func (i interface{}) bool {
		return i.(string)[:3] == "abc"
	}
	testIterEq(t, String("abcdef", "abcdaj"), TakeWhile(pred, Cycle(String("abcdef", "abcdaj", "ajcde"))))
}

func TestFilter(t *testing.T) {
	pred := func (i interface{}) bool {
		return i.(uint64) % 2 == 1
	}
	testIterEq(t, Uint64(1,3,5,7,9), Filter(pred, Uint64(1,2,3,4,5,6,7,8,9,10)))
	testIterEq(t, Uint64(2,4,6,8,10), FilterFalse(pred, Uint64(1,2,3,4,5,6,7,8,9,10)))
}

func TestSlice(t *testing.T) {
	testIter(t, Int(5,6,7,8,9,10), Slice(Count(0), 5))
	testIterEq(t, Int(2,3,4,5,6,7,8), Slice(Count(0), 2, 9))
	testIterEq(t, Int(3,6,9), Slice(Count(0), 3, 11, 3))
}

func TestMap(t *testing.T) {
	mapper := func (i interface{}) interface{} {
		return len(i.(string))
	}
	testIterEq(t, Int(1,2,3,4), Map(mapper, String("a", "ab", "abc", "abcd")))
}

func TestMultiMap(t *testing.T) {
	multiMapper := func (is ...interface{}) interface{} {
		var s float64
		for _, i := range is {
			s += i.(float64)
		}
		return s
	}
	testIterEq(t, Float64(10.4, 3.2), MultiMap(multiMapper, Float64(5.2, 1.6, 2.2), Float64(5.2, 1.0), Float64(0, 0.6, 0)))
}

func TestZip(t *testing.T) {
	a, b, c := []interface{}{1,"a"}, []interface{}{2,nil}, []interface{}{3,nil}
	test1, test2 := make(Iter, 1), make(Iter, 3)
	test1 <- a
	test2 <- a; test2 <- b; test2 <- c
	close(test1)
	close(test2)

	testIterEq(t, test1, Zip(Count(1), String("a")))
	testIterEq(t, test2, ZipLongest(Slice(Count(1), 0, 3), String("a")))
}

func TestStarmap(t *testing.T) {
	multiMapper := func (is ...interface{}) interface{} {
		var s int = 1
		for _, i := range is {
			s *= i.(int)
		}
		return s
	}
	testIterEq(t, Int(10, 20, 30), Starmap(multiMapper, Zip(Int(1,2,3), Repeat(10, 3))))
}

func TestReduce(t *testing.T) {
	summer := func (memo interface{}, el interface{}) interface{} {
		return memo.(float64) + el.(float64)
	}
	if float64(.82) - Reduce(Float64(.1,.2,.3,.22), summer, float64(0)).(float64) > .000001 {
		t.Error("Sum Reduce failed")
	}
}