golang:1.22-alpine3.18 AS build-env

RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make
RUN go install github.com/google/wire/cmd/wire@latest
WORKDIR /go/src/github.com/devtron-labs/telemetry
ADD . /go/src/github.com/devtron-labs/telemetry
RUN GOOS=linux make

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build-env  /go/src/github.com/devtron-labs/telemetry/telemetry .
CMD ["./telemetry"]