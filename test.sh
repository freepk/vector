GO111MODULE="off" go fmt
GO111MODULE="off" GOARCH="amd64" GOOS="linux" go test -v -gcflags="-B" -c -o test.amd64.linux
GO111MODULE="off" GOARCH="arm64" GOOS="darwin" go test -v -gcflags="-B" -c -o test.arm64.darwin
GO111MODULE="off" go test -v