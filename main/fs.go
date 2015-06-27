package main

import (
	"crypto/tls"
	"github.com/gregjones/httpcache"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"log"
	"net/http"
	"os"
)

type ProjectFs struct {
	pathfs.FileSystem
	ProjectStore ProjectStore
}

func NewProjectStore() ProjectStore {
	var tr *http.Transport
	if os.Getenv("UNSAFE_TLS") == "true" {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tr = &http.Transport{}
	}
	cachedTr := httpcache.NewMemoryCacheTransport()
	cachedTr.Transport = tr
	ps := ProjectStore{Transport: cachedTr}
	return ps
}

func MountDefaultProjectFs(path string) (*fuse.Server, error) {
	nfs := pathfs.NewPathNodeFs(&ProjectFs{
		FileSystem:   pathfs.NewDefaultFileSystem(),
		ProjectStore: NewProjectStore(),
	}, nil)
	server, _, err := nodefs.MountRoot(path, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
		return nil, err
	}
	return server, nil
}

func (me *ProjectFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	// stat a particular file
	for _, b := range me.ProjectStore.GetJsonIndex().Json {
		if b.Id == name {
			return &fuse.Attr{
				Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
			}, fuse.OK
		}
	}

	// stat the directory
	if name == "" {
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}

	return nil, fuse.ENOENT
}

func (me *ProjectFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	if name == "" {
		c := []fuse.DirEntry{}
		pji := me.ProjectStore.GetJsonIndex()
		for i := range pji.Json {
			log.Println(pji.Json[i].Id)
			c = append(c, fuse.DirEntry{Name: pji.Json[i].Id, Mode: fuse.S_IFREG})
		}

		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *ProjectFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	if flags&fuse.O_ANYWRITE != 0 {
		return nil, fuse.EPERM
	}

	body, err := me.ProjectStore.GetMarkdown(name)
	if err != nil {
		return nil, fuse.EPERM
	}

	return nodefs.NewDataFile([]byte(body)), fuse.OK
}
