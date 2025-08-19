# Build Stage
FROM golang:1.24.6-alpine AS build_stage
WORKDIR /app
COPY go.mod .
COPY go.sum .
# run `go mod vendor` on pipeline for a faster build
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/app ./cmd/main.go

# Deploy Stage
FROM alpine as deploy

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /app
COPY --from=build_stage /app/cmd/app .
COPY --from=build_stage /app/cmd/config.yaml ./cmd/
EXPOSE 8080
ENTRYPOINT ["./app"]