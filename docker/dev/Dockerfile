FROM golang:alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src

COPY . .

RUN apk --no-cache add curl

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-w -s" -o /go/bin/college ./cmd/main.go

CMD ["air", ".air.toml"]