git:
	git add .
	git commit -m "$(msg)"
	git push origin master

client:
	@echo "removing existing client binary"
	rm -f bin/client
	@echo "building client binary"
	go build -o bin/client cmd/client/main.go