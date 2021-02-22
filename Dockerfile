FROM golang:1.16-alpine AS builder
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /app/main .

FROM alpine:3.13
COPY --from=builder /dist/main /
CMD [ "/main" ]
