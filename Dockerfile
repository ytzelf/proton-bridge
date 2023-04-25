FROM golang:1.20

WORKDIR /opt/protonbridge

# Install dependencies
RUN apt-get update && apt-get install -y --no-install-recommends libglvnd-dev libsecret-1-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build-nogui

# CMD ["app"]
