proto:
	protoc --go_out=. ./proto/*.proto

serve:
	go run main.go -p 3031

tidy:
	go mod tidy

test:
	go run test 

build:
	go run build


.PHONY: proto serve tidy test build
