

build:
	go build urldecode.go
	go build urlencode.go


INSTALL_PATH = ~/.local/bin
install: build
	cp urldecode $(INSTALL_PATH)
	cp urlencode $(INSTALL_PATH)


.PHONY: test
test:
	cd test && make test
