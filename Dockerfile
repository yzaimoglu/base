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
# Now copy the static shell into base image.
COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/mkdir /bin/mkdir
COPY --from=busybox /bin/cat /bin/cat
COPY --from=busybox /bin/ls /bin/ls

COPY --from=build-stage /base/entrypoint /base/entrypoint
# Copy all static files from /ui/static to /ui/static in the container.
COPY --from=build-stage /app/ui/static /base/ui/static
ENTRYPOINT ["/base/entrypoint"]
CMD ["serve"]
