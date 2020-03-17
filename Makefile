export PATH := /Users/pquiring/bin/openapi-sdkgen-3.2.0:$(PATH)
PACKAGE=exampleservicev1

go:
	PATH=$(PATH) openapi-sdkgen.sh generate -g ibm-go -i ./exampleservice.yaml -o .
	rm -f go.mod go.sum
	go mod init github.com/powellquiring/my-sdk
	go mod tidy	

