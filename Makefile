
build:
	go run ./project-fs.go "`pwd`/mnt"

install:
	echo "No install script written :("

uninstall:

install-deps:
	export GOPATH=`pwd` && go get ./...

force-unmount:
	sudo umount -l "`pwd`/mnt"

