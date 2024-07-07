# STAGE 1
FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /goapi

# STAGE 2
FROM scratch

COPY --from=builder /goapi /goapi
ENV TZ=UTC

ENTRYPOINT ["/goapi"]