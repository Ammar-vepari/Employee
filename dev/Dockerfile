# --------------------------------------------------------------------- dev ---

    FROM golang:1.20.4-alpine as dev

    RUN apk --update add ca-certificates openssh-client alpine-sdk git

    ENV GO111MODULE=on \
        CGO_ENABLED=1 \
        # GOOS=linux \
        # GOARCH=amd64
        GOSUMDB=off \
        GOFLAGS=-buildvcs=false

    WORKDIR /go/src/github.com/services/employee

    COPY ./go.mod ./go.sum ./

    RUN go mod download

    COPY . .
    # ------------------------------------------------------------------- debug ---

    FROM dev as debug

    # install debugging tools
    RUN go install github.com/go-delve/delve/cmd/dlv@latest
    RUN go install github.com/cespare/reflex@latest

    # install any project dependencies
    RUN go get -d -v ./...

    CMD reflex -R "__debug_bin" -s -- sh -c "dlv debug --headless --continue --accept-multiclient --listen :40000 --build-flags -tags musl --api-version=2 --log ./cmd/"
