package sys

import (
	//"fmt"
	"runtime"
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

func CpuInject(stop *bool, wg *sync.WaitGroup) {

	log.Info("Sys CPU Inspector Start")

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(s *bool, w *sync.WaitGroup) {
			for {
				if *s {
					break
				}
				//if mprobab(1) {
				if probab(1) {
					//time.Sleep(1 * time.Millisecond)
					time.Sleep(1 * time.Microsecond)
				}
			}
			(*w).Done()
		}(stop, wg)
	}

	log.Info("Sys CPU Inspector Stop")
}
