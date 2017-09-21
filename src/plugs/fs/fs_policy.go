package fs

import (
	//"fmt"
	"sync"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"

	config "DistSysTestPlatform/src/utils/config"
	convert "DistSysTestPlatform/src/utils/convert"
	queue "DistSysTestPlatform/src/utils/queue"
	log "github.com/cihub/seelog"
)

func Policy(item_ interface{}) {
	var duration int
	var mountDir string
	var wg sync.WaitGroup

	item := item_.(*(queue.BasicTimeoutQueueItem))
	value := item.Value()

	switch value.(type) {
	case config.DiskConfig:
		_struct := value.(config.DiskConfig)
		//fmt.Println(_struct)
		p := _struct.Parameter
		duration = convert.String2Int(p.Duration)
		mountDir = p.Dir
		break
	default:
		return
		//break
	}

	//release mount point firstly
	fuse.ExecRelease(mountDir, "-l")
	time.Sleep(2 * time.Second)

	nfs := pathfs.NewPathNodeFs(&HelloFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(mountDir, nfs.Root(), nil)
	if err != nil {
		log.Error("Mount fail:\n", err)
		return
	}

	wg.Add(1)
	go func() {
		server.Serve()
		wg.Add(-1)
	}()

	if err := server.WaitMount(); err != nil {
		log.Error("WaitMountk:", err)
		return
	}

	go func() {
		timer1 := time.NewTicker(time.Duration(duration) * time.Second)
		<-timer1.C
		timer1.Stop()

		server.Unmount()
	}()

	wg.Wait()
	fuse.ExecRelease(mountDir, "-l")
	time.Sleep(2 * time.Second)
}
