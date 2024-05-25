FROM busybox:1.35.0-uclibc as busybox

# Build stage.
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /base/entrypoint

# Deploy stage.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /base
COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/mkdir /bin/mkdir
COPY --from=busybox /bin/cat /bin/cat
COPY --from=busybox /bin/ls /bin/ls

COPY --from=build-stage /base/entrypoint /base/entrypoint
COPY --from=build-stage /app/ui/static /base/ui/static
COPY --from=build-stage /app/migrations /base/migrations
ENTRYPOINT ["/base/entrypoint"]
CMD ["serve"]
