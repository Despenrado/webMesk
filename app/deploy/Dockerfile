FROM golang:1.18 AS builder
WORKDIR /go_rest
COPY . .
# RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o build/restapi ./cmd/main.go

FROM scratch
EXPOSE 8081
WORKDIR /go_rest
COPY ./configs/backend.yaml .
COPY --from=builder /go_rest/build .
ENTRYPOINT ["./restapi","-config=./backend.yaml"]
