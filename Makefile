lint:
	golint -set_exit_status
	megacheck

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic
