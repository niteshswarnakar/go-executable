build:
	go build -o extension
	./zipper

build-windows:
	go build -o extension.exe
	.\zipper.exe