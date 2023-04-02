

build:
	go build urldecode.go
	go build urlencode.go


INSTALL_PATH = /home/zrth/.local/bin
install:
	cp urldecode $(INSTALL_PATH)
	cp urlencode $(INSTALL_PATH)
