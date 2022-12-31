FROM golang:1.19-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Disable cgo, set os and platfrom
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Strip DWARF, symbol table and debug info to decrease binary size
# https://pkg.go.dev/cmd/link
RUN go build -ldflags="-s -w" -o insurance-api .

FROM scratch
COPY --from=builder ["/build/insurance-api", "/build/.env", "/"]
EXPOSE 8080
ENTRYPOINT [ "/insurance-api" ]