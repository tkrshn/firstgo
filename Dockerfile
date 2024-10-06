# syntax=docker/dockerfile:1

FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy the whole directory into /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-app

# Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM debian:latest AS build-release-stage

WORKDIR /

ARG GH_TOKEN
ENV TOKEN=$GH_TOKEN

# display env TOKEN how to run this command
RUN echo $GH_TOKEN
RUN echo $TOKEN

COPY --from=build-stage /docker-go-app /docker-go-app

ENTRYPOINT ["/docker-go-app"]