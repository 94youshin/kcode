build:
	go build  -o _output/bin/concli con-cri.go 
debug: build
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./con-cri
