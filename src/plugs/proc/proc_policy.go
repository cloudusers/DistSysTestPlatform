package proc

import (
	"fmt"
	"math/rand"
	"time"

	config "DistSysTestPlatform/src/utils/config"
	convert "DistSysTestPlatform/src/utils/convert"
	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func Policy(item_ interface{}) {

	//var duration int = 0
	var signal int = 0
	//var stop bool = false
	item := item_.(*(queue.BasicTimeoutQueueItem))
	value := item.Value()

	switch value.(type) {
	case config.ProcConfig:
		_struct := value.(config.ProcConfig)
		p := _struct.Parameter
		signal = convert.String2Int(p.Signal)
		//duration = convert.String2Int(p.Duration)

		break
	default:
		return
	}

	//timer
	/*
		go func(stop *bool) {
			timer1 := time.NewTicker(time.Duration(duration) * time.Second)
			<-timer1.C
			timer1.Stop()
			*stop = true
		}(&stop)
	*/

	for {
		/*
			if stop {
				break
			}
		*/

		if 0 == signal {
		} else if 15 == signal {
		} else if 9 == signal {
		} else {
			msg := fmt.Sprintf("unknown signal:%d", signal)
			log.Debug(msg)
		}

		time.Sleep(1 * time.Second)
		break
	}
}

func probab(percentage int) bool {
	return rand.Intn(99) < percentage
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
