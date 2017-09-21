package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type HelloFs struct {
	pathfs.FileSystem
}

func (me *HelloFs) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	data := []byte{'d', 's', 't', 'p'}
	return fuse.ReadResultData(data), fuse.OK
}

func (me *HelloFs) Write(d []byte, off int64) (uint32, fuse.Status) {
	return uint32(16), fuse.OK
}
func (me *HelloFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	/*
		switch name {
		case "dstp.txt":
			return &fuse.Attr{
				//Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
				Mode: fuse.S_IFREG | 0666, Size: uint64(len(name)),
			}, fuse.OK

		case "":
			return &fuse.Attr{
				Mode: fuse.S_IFDIR | 0755,
			}, fuse.OK
		}
		return nil, fuse.ENOENT
	*/
	switch name {
	case "":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	default:
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
		}, fuse.OK
	}
}

func (me *HelloFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	if name == "" {
		c = []fuse.DirEntry{{Name: "dstp.txt", Mode: fuse.S_IFREG}}
		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *HelloFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	/*
		if name != "dstp.txt" {
			return nil, fuse.ENOENT
		}
		if flags&fuse.O_ANYWRITE != 0 {
			return nil, fuse.EPERM
		}
	*/
	return nodefs.NewDataFile([]byte(name)), fuse.OK
}

func (me *HelloFs) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.OK
}
