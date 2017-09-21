package fs

import (
	"time"

	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func Run() {
	log.Info("Disk Inspector Start")
	for {
		select {
		case deq := <-queue.QueueFactory.DiskQueueCh:
			Policy(deq)
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Info("Disk Inspector Stop")
}
