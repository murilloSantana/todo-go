build: 
	go build -v -o bin/todo-go

test:
	go test -v --cover ./...

run:
	docker-compose up --build