FROM golang:1.14.9-alpine3.12
WORKDIR /usr/app
COPY . .
RUN go build -o bin/todo-go
EXPOSE 8080
CMD [ "./bin/todo-go" ]