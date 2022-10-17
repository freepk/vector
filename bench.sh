GO111MODULE="off" go fmt
GO111MODULE="off" go test -v -bench . -benchmem -gcflags="-B"
