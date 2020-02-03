FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .

RUN go mod download

COPY main.go                .

RUN ls /app/
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o my-go-extension .

FROM scratch
WORKDIR /app
COPY --from=builder /app/my-go-extension /app/

EXPOSE 8080
ENTRYPOINT ["/app/my-go-extension"]