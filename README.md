# GoLang features
In this project I collect some of the features which I've tried

## Cross compile
### Before compile
1. Make sure that you've installed the docker on your PC
2. Make sure that you are inside of GOPATH directory
3. Make sure that you installed all the vendors by running `make vendor` command
4. Make sure that you pull the docker image which will be used during compilation `docker pull karalabe/xgo-latest`
5. Make sure that you download the xgo vendor `go get github.com/karalabe/xgo`

### Compile with CGO enabled
To compile the test project please run the following command: `make build-cgo-enabled` 

### Compile with CGO disabled
To compile the test project please run the following command: `make build` 
