package utils

type Intset map[int]struct{}

func (s Intset) Add(value int) {
	s[value] = struct{}{}
}

func (s Intset) Contains(value int) bool {
	_, ok := s[value]
	return ok
}

func (s Intset) Union(other Intset) Intset {
	combined := Intset{}
	for id := range s {
		combined.Add(id)
	}
	for id := range other {
		combined.Add(id)
	}
	return combined
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

func Square(n int) int {
	return n * n
}
