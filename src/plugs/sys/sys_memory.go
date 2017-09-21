package sys

import (
	"container/list"
	"math/rand"
	//"runtime"
	"time"

	log "github.com/cihub/seelog"
)

var makes int
var frees int

func makeBuffer() []byte {
	makes += 1
	return make([]byte, rand.Intn(900000000)+900000000)
}

type queued struct {
	when  time.Time
	slice []byte
}

func makeRecycler() (get, give chan []byte) {
	get = make(chan []byte)
	give = make(chan []byte)

	go func() {
		q := new(list.List)
		for {
			if q.Len() == 0 {
				q.PushFront(queued{when: time.Now(), slice: makeBuffer()})
			}

			e := q.Front()

			timeout := time.NewTimer(time.Minute)
			select {
			case b := <-give:
				timeout.Stop()
				q.PushFront(queued{when: time.Now(), slice: b})

			case get <- e.Value.(queued).slice:
				timeout.Stop()
				q.Remove(e)

			case <-timeout.C:
				e := q.Front()
				for e != nil {
					n := e.Next()
					if time.Since(e.Value.(queued).when) > time.Minute {
						q.Remove(e)
						e.Value = nil
					}
					e = n
				}
			}
		}

	}()

	return
}

func MemInject(stop *bool) {

	log.Info("Sys MEM Inspector Start")

	func(s *bool) {
		pool := make([][]byte, 20)
		get, give := makeRecycler()

		//var m runtime.MemStats
		for {
			if *s {
				break
			}

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

			/*
				runtime.ReadMemStats(&m)
				fmt.Printf("heapsys:%d, bytes:%d, headalloc:%d, headidle:%d, headreleased:%d, makes:%d, frees:%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased, makes, frees)
			*/
		}
	}(stop)

	log.Info("Sys MEM Inspector Stop")
}
