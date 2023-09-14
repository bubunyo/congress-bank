package congressbank

import "math/rand"

var (
	min = 1_000_000_000
	max = 1_000_000_0000 - 1
)

func genId() int {
	r1 := rand.New(src)
	return r1.Intn(max-min) + min
}
