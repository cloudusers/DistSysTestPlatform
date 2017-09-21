package ethernet

import (
	"time"

	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func RunTc() {
	log.Info("Ethernet TC Inspector Start")

	for {
		select {
		case deq := <-queue.QueueFactory.NetQueueCh:
			PolicyTc(deq)
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}

	log.Info("Ethernet TC Inspector Stop")
}
