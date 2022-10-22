GO111MODULE="off" go fmt
GO111MODULE="off" GOARCH="amd64" go test -v -bench . -benchmem -gcflags="-B"
