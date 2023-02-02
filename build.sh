CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o ./build/go-monitor-linux-amd64

CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build -o ./build/go-monitor-windows-amd64
