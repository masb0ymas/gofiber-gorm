ARG APP_NAME=my-app

FROM golang:1.20-alpine as base
LABEL author="masb0ymas"
LABEL name="expresso"

FROM base as build_base

RUN apk add --no-cache git

# Set the Temp Working Directory inside the container
WORKDIR /temp-build

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cp .env.docker-production .env

# Build the Go app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o ./dist/${APP_NAME} main.go

# Start fresh from a smaller image
FROM alpine:3.16 as runner
RUN apk add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Setup Timezone
RUN apk add tzdata
ENV TZ=Asia/Jakarta

# editor cli with nano
RUN apk add nano

COPY --from=build_base /temp-build/dist/${APP_NAME} ./${APP_NAME}
COPY --from=build_base /temp-build/bin ./bin
COPY --from=build_base /temp-build/public ./public
COPY --from=build_base /temp-build/.air.toml ./.air.toml
COPY --from=build_base /temp-build/.env ./.env

# This container exposes port 8000 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["./${APP_NAME}"]
