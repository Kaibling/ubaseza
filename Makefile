build:
	go mod tidy && go build -o zucktrans
start:
	./zucktrans
test:
	go test -v ./... -race -covermode=atomic