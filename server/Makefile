.PHONY: build-server build-client clean

build-server:
	go build -o server/server server/src/main.go

build-client:
	cd client/tourist_client && make

clean:
	rm -rf server/server client/tourist_client/*.o
