build:
	go build -o bin/aip2p aip2p.go

deb-package:
    fyne package -os linux
	tar -C ./aip2p-deb -xf AIP2P.tar.xz
	rm ./aip2p-deb/Makefile
    dpkg-deb --build weblang-deb