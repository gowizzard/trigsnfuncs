# Dockerfile for the trigsnfuncs application. This file is used to build the production-ready image of the application.
# The image is based on the Alpine Linux distribution, which is known for its small size and security.

########################################################################################################################

# This phase uses the Alpine-based Go image to compile the source code of the application.
# By parameterizing the Go version, it becomes straightforward to maintain and modify in the future.
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --no-cache git make
WORKDIR /tmp/src
COPY . .
RUN make build

########################################################################################################################

# The final preparation phase for the production-ready image. Essential system packages are added,
# a dedicated user for the application is created (enhancing security by avoiding root privileges),
# and the correct timezone is set.
FROM alpine:latest AS production
RUN apk add --no-cache tzdata && \
    adduser -D application
USER application
ENV TZ=Europe/Berlin
WORKDIR /app
COPY --from=build /tmp/src .
ENTRYPOINT ["/app/trigsnfuncs"]