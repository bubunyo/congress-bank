package congressbank

import (
	"fmt"
	"math/rand"
)

var (
	min = 1_000_000_000
	max = 1_000_000_0000 - 1
)

func genId() int {
	r1 := rand.New(src)
	return r1.Intn(max-min) + min
}

func genIdWithPrefix(p string) string {
	r1 := rand.New(src)
	id := r1.Intn(max-min) + min
	return fmt.Sprintf("%s%d", p, id)

}
