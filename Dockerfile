FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server ./cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/server /app

EXPOSE 8080

CMD ["./server"]
