build:
	CGO_ENABLED=0 GOOS=linux go build -o dist/whereami -a -tags netgo -ldflags '-w' .
	docker build -t whereami .
