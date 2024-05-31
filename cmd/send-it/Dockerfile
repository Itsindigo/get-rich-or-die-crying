FROM golang:1.22

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY ./cmd/send-it/*.go ./cmd/send-it/
COPY ./internal ./internal

RUN go build -o /app/cmd/send-it ./cmd/send-it

RUN chmod +x /app/cmd/send-it

ENTRYPOINT ["/app/cmd/send-it/send-it"]
