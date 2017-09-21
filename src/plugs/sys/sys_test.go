package sys

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"testing"
)

func TestMain(m *testing.M) {
	pool := make([][]byte, 20)

	get, give := makeRecycler()

	var m runtime.MemStats
	for {
		b := <-get
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			give <- pool[i]
		}

		pool[i] = b

		time.Sleep(100 * time.Millisecond)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("heapsys:%d, bytes:%d, headalloc:%d, headidle:%d, headreleased:%d, makes:%d, frees:%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased, makes, frees)
	}
}
