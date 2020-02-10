LDFLAGS=-ldflags="-s -w"
BIN_DIR_CGO_DISABLED=bin/cgo-disabled
BIN_DIR_CGO_ENABLED=bin/cgo-disabled

vendor:
	if [ ! -d "vendor" ] || [ -z "$(shell ls -A vendor)" ]; then go mod vendor; fi

build-cgo-enabled:
	env CGO_ENABLED=1 xgo --targets=darwin/*,linux/amd64,linux/386,windows/* --dest ./$(BIN_DIR_CGO_ENABLED) --out build ./cross-compile

build: cross-compile/main.go vendor
	env GOOS=darwin go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-mac $<
	env GOOS=linux GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-linux-amd64 $<
	env GOOS=linux GOARCH=386 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-linux-386 $<
	env GOOS=freebsd GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-freebsd-amd64 $<
	env GOOS=freebsd GOARCH=386 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-freebsd-386 $<
	env GOOS=windows GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-windows-amd64.exe $<
	env GOOS=windows GOARCH=386 go build -mod=vendor $(LDFLAGS) -o $(BIN_DIR_CGO_DISABLED)/$@-windows-386.exe $<

.PHONY: vendor build