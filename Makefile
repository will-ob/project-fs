
# Where GOPATH is an env var ~/projects/go-workspace
build:
	cd "$$GOPATH" && go run src/github.com/will-ob/todo/main/*go src/github.com/will-ob/todo/mnt

install:
	echo "No install script written :("

uninstall:

install-deps:
	export GOPATH=`pwd` && go get ./...

force-unmount:
	cd "$$GOPATH" && sudo umount -l src/github.com/will-ob/todo/mnt

