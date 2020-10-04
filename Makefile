build: 
	go build -v -o bin/todo-go

test:
	go test -v --cover ./...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
run:
	docker-compose up --build