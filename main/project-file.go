package main

import (
	"bytes"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"log"
	"time"
)

type projectFile struct {
	data  []byte
	store *ProjectStore
	name  string
	nodefs.File
}

func NewProjectFile(data []byte, ps *ProjectStore, name string) nodefs.File {
	f := new(projectFile)
	f.data = data
	f.store = ps
	f.name = name
	f.File = nodefs.NewDefaultFile()
	return f
}

func (f *projectFile) SetInode(*nodefs.Inode) {
}

func (f *projectFile) InnerFile() nodefs.File {
	return nil
}

func (f *projectFile) String() string {
	return "projectFile"
}

func (f *projectFile) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	log.Println("Reading file...")
	end := int(off) + int(len(buf))
	if end > len(f.data) {
		end = len(f.data)
	}

	return fuse.ReadResultData(f.data[off:end]), fuse.OK
}

func (f *projectFile) Write(data []byte, off int64) (uint32, fuse.Status) {
	log.Println("Writing file...")
	log.Println(data)
	var err error
	_, err = f.store.SetMarkdown(f.name, &data)
	if err != nil {
		return 0, fuse.ENOENT
	} else {
		size := bytes.NewReader(data).Len()
		return uint32(size), fuse.OK
	}
}

func (f *projectFile) Flush() fuse.Status {
	return fuse.OK
}

func (f *projectFile) Release() {

}

func (f *projectFile) GetAttr(out *fuse.Attr) fuse.Status {
	out.Mode = fuse.S_IFREG | 0644
	out.Size = uint64(len(f.data))

	return fuse.OK
}

func (f *projectFile) Fsync(flags int) (code fuse.Status) {
	return fuse.OK
}

func (f *projectFile) Utimens(atime *time.Time, mtime *time.Time) fuse.Status {
	return fuse.ENOSYS
}

func (f *projectFile) Truncate(size uint64) fuse.Status {
	return fuse.OK
}

func (f *projectFile) Chown(uid uint32, gid uint32) fuse.Status {
	return fuse.ENOSYS
}

func (f *projectFile) Chmod(perms uint32) fuse.Status {
	return fuse.ENOSYS
}

func (f *projectFile) Allocate(off uint64, size uint64, mode uint32) (code fuse.Status) {
	return fuse.OK
}
