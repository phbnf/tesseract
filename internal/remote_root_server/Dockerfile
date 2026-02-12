FROM golang:1.24.5-alpine3.21@sha256:933e5a0829a1f97cc99917e3b799ebe450af30236f5a023a3583d26b5ef9166f AS builder

ARG GOFLAGS="-trimpath -buildvcs=false -buildmode=exe"
ENV GOFLAGS=$GOFLAGS

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/remote_root_server ./internal/remote_root_server

FROM alpine:3.22.0@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715

COPY --from=builder /build/bin/remote_root_server /bin/remote_root_server
ENTRYPOINT ["/bin/remote_root_server"]
