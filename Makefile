git:
	git add .
	git commit -m "$(msg)"
	git push origin master

client:
	@echo "removing existing client binary"
	rm -f bin/client
	@echo "building client binary"
	go build -o bin/client cmd/client/main.go

server:
	@echo "removing existing server binary"
	rm -f bin/server
	@echo "building server binary"
	go build -o bin/server cmd/server/main.go

fmt:
	@echo "formatting code"
	go fmt ./...

lint:
	@echo "Linting the source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...


.PHONY: client server