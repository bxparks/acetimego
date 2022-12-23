all:
	go build epoch.go

test:
	go test epoch.go epoch_test.go
