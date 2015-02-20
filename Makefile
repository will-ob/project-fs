
build:
	echo "Nothing to bulid"

install:
	echo "No install script written :("

uninstall:

install-deps:
	export GOPATH=`pwd` && go get ./...

force-unmount:
	# sudo umount -l <path>

