current_dir=$(shell pwd)

build: clean
	go build -o target/projectfs ./main/*.go

clean:
	rm -rf target

install: uninstall build
	$(current_dir)/tools/install.sh

uninstall:
	$(current_dir)/tools/uninstall.sh

install-deps:
	export GOPATH=`pwd` && go get ./...

force-unmount:
	sudo umount -l ./mnt

.PHONY: build clean install uninstall install-deps force-unmount
