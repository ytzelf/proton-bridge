FROM golang:1.20

WORKDIR /opt/protonbridge

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build-nogui

# CMD ["app"]
