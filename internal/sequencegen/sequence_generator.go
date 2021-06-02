// Package sequencegen provides functions for generating bool sequences.
package sequencegen

// makeInitialW makes the initial sequence for the Weight function.
func makeInitialW(n, w int) []bool {
	b := make([]bool, n)
	// First element is 01^w0^n-w-1
	for i := 1; i <= w; i++ {
		b[i] = true
	}
	return b
}

// Weight returns a closure that acts as an iterator of binary sequences of size
// n and weight w.
func Weight(n, w int) func() []bool {
	if w == 0 {
		return func() []bool {
			return make([]bool, n)
		}
	}
	if w == n {
		return func() []bool {
			r := make([]bool, n)
			for i := 0; i < n; i++ {
				r[i] = true
			}
			return r
		}
	}
	init := true
	x := 1
	y := 0
	b := makeInitialW(n, w)
	return func() []bool {
		if init {
			init = false
			return b
		}
		if x < n {
			b[x] = false
			b[y] = true
			x++
			y++
			if x < n && b[x] == false {
				b[x] = true
				b[0] = false
				if y > 1 {
					x = 1
				}
				y = 0
			}
			return b
		} else {
			b = makeInitialW(n, w)
			x = 1
			y = 0
			return b
		}
	}
}

func makeInitialB(n int) []bool {
	b := make([]bool, n)
	// First element is 01^n-1
	for i := 1; i < n; i++ {
		b[i] = true
	}
	return b
}

// Binary returns a closure that acts as an iterator of binary sequences of size
// n.
func Binary(n int) func() []bool {
	init := true
	x := 1
	y := 0
	b := makeInitialB(n)
	return func() []bool {
		if init {
			init = false
			return b
		}
		b2 := make([]bool, n)
		copy(b2, b)
		b = b2
		if y < n {
			b[x] = false
			b[y] = true
			x++
			y++
			if x == n && y != n {
				b[0] = false
				y = 0
				var bit int
				if b[1] {
					bit = 1
				}
				x = 0 + bit
			} else if x != y && !b[x] {
				b[x] = true
				b[0] = false
			}
			return b
		} else {
			b = makeInitialB(n)
			x = 1
			y = 0
			return b
		}
	}
}
