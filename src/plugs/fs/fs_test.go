package fs

import (
	"flag"
	"fmt"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/osrg/hookfs/hookfs"
	"github.com/osrg/namazu/nmz/endpoint/local"
	"github.com/osrg/namazu/nmz/signal"
	logutil "github.com/osrg/namazu/nmz/util/log"
	"github.com/osrg/namazu/nmz/util/mockorchestrator"
	"github.com/stretchr/testify/assert"
	//	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	flag.Parse()
	logutil.InitLog("", true)
	signal.RegisterKnownSignals()
	orcActionCh := make(chan signal.Action)
	orcEventCh := local.SingletonLocalEndpoint.Start(orcActionCh)
	defer local.SingletonLocalEndpoint.Shutdown()
	mockOrc := mockorchestrator.NewMockOrchestrator(orcEventCh, orcActionCh)
	mockOrc.Start()
	defer mockOrc.Shutdown()
	os.Exit(m.Run())
}

/*
func TestFilesystemInspectorInterfaceImpl(t *testing.T) {
	h := &FilesystemInspector{}
	func(x hookfs.HookWithInit) {}(h)
	func(x hookfs.HookOnRead) {}(h)
	func(x hookfs.HookOnWrite) {}(h)
	func(x hookfs.HookOnMkdir) {}(h)
	func(x hookfs.HookOnRmdir) {}(h)
	func(x hookfs.HookOnOpenDir) {}(h)
	func(x hookfs.HookOnFsync) {}(h)
}
*/

func newFUSEServer(t *testing.T, fs *hookfs.HookFs) *fuse.Server {
	opts := &nodefs.Options{
		NegativeTimeout: time.Second,
		AttrTimeout:     time.Second,
		EntryTimeout:    time.Second,
	}
	pathFs := pathfs.NewPathNodeFs(fs, nil)
	conn := nodefs.NewFileSystemConnector(pathFs.Root(), opts)
	originalAbs, _ := filepath.Abs(fs.Original)
	options := []string{"nonempty"}
	mOpts := &fuse.MountOptions{
		AllowOther: false,
		Name:       fs.FsName,
		FsName:     originalAbs,
		Options:    options,
	}
	server, err := fuse.NewServer(conn.RawFS(), fs.Mountpoint, mOpts)
	assert.NoError(t, err)
	server.SetDebug(true)
	return server
}

func TestFilesystemInspector_3(t *testing.T) {
	testFilesystemInspector(t, 3)
}

func testFilesystemInspector(t *testing.T, n int) {
	insp := FilesystemInspector{
		OrchestratorURL: "local://",
		EntityID:        "dummy",
	}

	origDir := `/data/logs/mount`
	mountDir := `/data/logs/mount`
	//defer os.RemoveAll(origDir)
	//mountDir, err := ioutil.TempDir("/tmp/aaa", "fs-test-mount")

	fs, err := hookfs.NewHookFs(origDir, mountDir, &insp)
	assert.NoError(t, err)
	fuseServer := newFUSEServer(t, fs)
	go fuseServer.Serve()
	fuseServer.WaitMount()
	defer fuseServer.Unmount()

	for i := 0; i < 30; i++ {
		fmt.Printf("---------%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
