.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/whois whois/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/lookupHost lookupHost/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/lookupAddr lookupAddr/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/lookupMX lookupMX/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/lookupNS lookupNS/main.go
	
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
