test:
	go test ./... -cover       

test-cover:
	go test ./... -covermode=atomic -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	
cover-strx:
	go test ./strx/... -cover