
gobuild = mkdir -p dist/$$GOOS-$$GOARCH && go build -o dist/$$GOOS-$$GOARCH

build: 
	export GOOS=linux; export GOARCH=386; $(gobuild)
	export GOOS=linux; export GOARCH=amd64; $(gobuild)
	export GOOS=windows; export GOARCH=amd64; $(gobuild)
	export GOOS=windows; export GOARCH=386; $(gobuild) 

test:
	go test ./...

clean:
	rm -rf dist

# Resources:
# List of available target OSs and architectures: 
# https://golang.org/doc/install/source#environment