############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS build

RUN apk update && apk add --no-cache git bash

WORKDIR $GOPATH/src

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-w -s" -o /go/bin/college ./cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=build /go/bin/college /go/bin/college

ENTRYPOINT [ "/go/bin/college" ]