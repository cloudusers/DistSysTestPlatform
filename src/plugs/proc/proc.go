package proc

import (
	"time"

	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func Run() {
	log.Info("Proc Inspector Start")
	for {
		select {
		case deq := <-queue.QueueFactory.ProcQueueCh:
			Policy(deq)
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Info("Proc Inspector Stop")
}
