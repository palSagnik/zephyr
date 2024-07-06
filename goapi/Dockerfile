FROM golang:1.21

# The app directory to hold the source code
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /goapi
# GIN_MODE=release use when in production
ENV TZ=UTC

CMD ["/goapi"]