FROM golang:1.18

WORKDIR /usr/src/logistics

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./cmd/...  # build all main.go files in ./cmd folder
