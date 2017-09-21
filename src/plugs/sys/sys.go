package sys

import (
	"time"

	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func RunCpu() {
	log.Info("Syscpu Inspector Start")
	for {
		select {
		case deq := <-queue.QueueFactory.SysCpuQueueCh:
			PolicyCpu(deq)
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}

	log.Info("Syscpu Inspector Stop")
}

func RunMem() {
	log.Info("Sysmem Inspector Start")
	for {
		select {
		case deq := <-queue.QueueFactory.SysMemQueueCh:
			PolicyMem(deq)
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}

	log.Info("Sysmem Inspector Stop")
}
