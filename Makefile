build:
	cd src && go build -o ../cfr

exec:
	$(MAKE) build
	./cfr 1000000

test:
	go test -v ./...

debug:
	dlv debug src/main.go
