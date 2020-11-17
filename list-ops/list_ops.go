// Package listops contains various functions for working with lists of integers.
package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append adds all items of the second list to the end of the first.
func (l IntList) Append(l2 IntList) IntList {
    for _, v := range l2 {
        l = append(l, v)
    }
    return l
}

// Concat combines all lists into one flattened list.
func (l IntList) Concat(listSlice []IntList) IntList {
    for _, list := range listSlice {
        for _, v := range list {
            l = append(l, v)
        }
    }
    return l
}

// Filter returns a list of all items for with the input predicate function returns true.
func (l IntList) Filter(f func(int) bool) IntList {
    out := IntList{}
    for _, v := range l {
		if f(v) {
            out = append(out, v)
		}
    }
    return out
}

// Length returns the total number of items within a list.
func (l IntList) Length() int {
    return len(l)
}

// Map returns a list of the results of applying the input function against all items.
func (l IntList) Map(f func(int) int) IntList {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}

// Foldl reduces each item into the accumulator from the left using the input function.
func (l IntList) Foldl(f binFunc, a int) int {
	for _, v := range l {
		a = f(a, v)
	}
	return a
}

// Foldr reduces each item into the accumulator from the right using the input function.
func (l IntList) Foldr(f binFunc, a int) int {
	for i := len(l) - 1; i >= 0; i-- {
		a = f(l[i], a)
	}
	return a
}

// Reverse returns a list with all the original items, but in reversed order.
func (l IntList) Reverse() IntList {
    out := IntList{}
    for i := len(l) - 1; i >= 0; i-- {
        out = append(out, l[i])
    }
    return out
}
