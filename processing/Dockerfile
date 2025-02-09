FROM golang:1.16-alpine AS build

# gcc and libc (musl-dev) are required by cryptixd
# git is required to build cryptixd by commit-hash
RUN apk add gcc musl-dev git

RUN mkdir /build
WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .

ARG CRYPTIXD_VERSION

# Exit if CRYPTIXD_VERSION is not set
RUN if [ -z "${CRYPTIXD_VERSION}" ]; then exit 1; fi
ARG CRYPTIXD_REPOSITORY='github.com/cryptix-network/cryptixd'

# Replace the cryptixd version in go.mod
RUN go mod edit -replace github.com/cryptix-network/cryptixd=$CRYPTIXD_REPOSITORY@$CRYPTIXD_VERSION

RUN go mod download

RUN cp go.mod go.mod.bu
RUN cp go.sum go.sum.bu

COPY . .

# Restore go.mod and go.sum because `COPY . .` overwrote them
RUN mv go.mod.bu go.mod
RUN mv go.sum.bu go.sum
RUN go mod tidy

RUN go env -w GOFLAGS=-mod=mod

RUN GOOS=linux go build -o processing .

FROM alpine
WORKDIR /app
COPY --from=build /build/processing /app/

RUN mkdir -p /app/database/migrations
COPY --from=build /build/database/migrations/ /app/database/migrations/
