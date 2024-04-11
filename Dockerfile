# Static files.
FROM node:20.11.1-alpine as frontend-builder
WORKDIR /builder
COPY /ui/package.json /ui/package-lock.json ./
ENV PUBLIC_API_URL="/api/v1/"
ENV PUBLIC_POCKETBASE_URL="/"
RUN npm ci
COPY /ui .
RUN npm run build

# Build.
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
COPY --from=frontend-builder /builder/build ./ui/build/
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
