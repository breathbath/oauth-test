# -------------
# Build stage

FROM golang:1.15 AS builder

ENV APP_DIR=/app

RUN mkdir -p $APP_DIR

RUN groupadd -g 1000 t1 && useradd -r -s /bin/false -u 1000 -g t1 t1

WORKDIR $APP_DIR
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $APP_DIR/auth main.go

# -------------
# Image creation stage

FROM alpine:latest
ENV APP_DIR=/app
WORKDIR $APP_DIR
COPY --from=builder /app/auth $APP_DIR/auth
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

RUN chown -R 1000:t1 $APP_DIR
USER 1000

CMD $APP_DIR/auth
