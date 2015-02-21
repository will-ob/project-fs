package main

import (
	"encoding/json"
	"flag"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type ProjectFs struct {
	pathfs.FileSystem
}

type test_struct struct {
	Id string
}

func (me *ProjectFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	resp, err := http.Get("http://localhost:3333/projects/")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var t []test_struct
	errr := decoder.Decode(&t)
	if errr != nil {
		log.Fatal(errr)
	}
	for _, b := range t {
		if b.Id == name {
			return &fuse.Attr{
				Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
			}, fuse.OK
		}
	}

	switch name {
	case "file.txt":
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
		}, fuse.OK
	case "":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *ProjectFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	resp, err := http.Get("http://localhost:3333/projects/")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var t []test_struct
	errr := decoder.Decode(&t)
	if errr != nil {
		log.Fatal(errr)
	}

	if name == "" {
		c := []fuse.DirEntry{}
		for i := range t {
			log.Println(t[i].Id)
			c = append(c, fuse.DirEntry{Name: t[i].Id, Mode: fuse.S_IFREG})
		}

		// {{Name: "file.txt", Mode: fuse.S_IFREG}}
		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *ProjectFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	if flags&fuse.O_ANYWRITE != 0 {
		return nil, fuse.EPERM
	}

	resp, err := http.Get("http://localhost:3333/projects/" + name)
	if err != nil {
		return nil, fuse.EPERM
	}

	defer resp.Body.Close()
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		return nil, fuse.EPERM
	}

	return nodefs.NewDataFile([]byte(body)), fuse.OK
}

func main() {

	// Check args
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  project-fs MOUNTPOINT")
	}

	// Load file system
	log.Println("Loading file system...")
	nfs := pathfs.NewPathNodeFs(&ProjectFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}

	go func() {
		server.Serve()
	}()
	log.Println("File system loaded.")

	// Clean 'errything up when I get SIGINT'd
	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c

		log.Printf("Captured %v\n", sig)
		log.Println("Unmounting file system...")

		server.Unmount()

		log.Println("File system unmounted.")
		done <- true
	}()
	<-done
}
