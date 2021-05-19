// Package sequencegen provides functions for generating bool sequences.
package sequencegen

// makeInitial makes the initial sequence for the Weight function.
func makeInitial(n, w int) []bool {
	b := make([]bool, n)
	// First element is 01^w0^n-w-1
	for i := 1; i <= w; i++ {
		b[i] = true
	}
	return b
}

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
	b := makeInitial(n, w)
	return func() []bool {
		if init {
			init = false
			return b
		}
		b2 := make([]bool, n)
		copy(b2, b)
		b = b2
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
			b = makeInitial(n, w)
			x = 1
			y = 0
			return b
		}
	}
}
