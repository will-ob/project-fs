package main

import (
	"flag"
	"runtime"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Check args
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  project-fs MOUNTPOINT")
	}

	// Load file system
	log.Println("Loading file system...")
	server, err := MountDefaultProjectFs(flag.Arg(0))
	if err != nil {
		log.Fatal("Failed to mount default project file system!")
		return
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
