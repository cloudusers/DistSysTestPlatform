package sys

import (
	//"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	config "DistSysTestPlatform/src/utils/config"
	convert "DistSysTestPlatform/src/utils/convert"
	queue "DistSysTestPlatform/src/utils/queue"
)

func PolicyCpu(item_ interface{}) {

	var duration int = 0
	var stop bool = false
	item := item_.(*(queue.BasicTimeoutQueueItem))
	value := item.Value()

	switch value.(type) {
	case config.SysCpuConfig:
		_struct := value.(config.SysCpuConfig)
		p := _struct.Parameter
		duration = convert.String2Int(p.Duration)
		break
	default:
		return
	}

	//
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	var cpus int = int(runtime.NumCPU())

	timer1 := time.NewTicker(time.Duration(duration) * time.Second)
	if probab(100) {
		wg.Add(cpus)
		CpuInject(&stop, &wg)
	}

	//
	<-timer1.C
	timer1.Stop()
	stop = true
	wg.Wait()
}

func PolicyMem(item_ interface{}) {

	var duration int = 0
	var stop bool = false
	item := item_.(*(queue.BasicTimeoutQueueItem))
	value := item.Value()

	switch value.(type) {
	case config.SysMemConfig:
		_struct := value.(config.SysMemConfig)
		p := _struct.Parameter
		duration = convert.String2Int(p.Duration)
		break
	default:
		return
	}

	//
	timer1 := time.NewTicker(time.Duration(duration) * time.Second)
	if probab(70) {
		MemInject(&stop)
	}

	//
	<-timer1.C
	timer1.Stop()
	stop = true
}

func probab(percentage int) bool {
	return rand.Intn(99) < percentage
}

/*
func mprobab(percentage int) bool {
	r := rand.Int63n(9999999)
	return r < int64(percentage)
	//return rand.Int63n(9999999) < percentage
}
*/

func init() {
	rand.Seed(time.Now().UnixNano())
}
