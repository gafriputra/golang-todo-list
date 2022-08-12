# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o /product-service

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /product-service /product-service
COPY .env /

EXPOSE 3030

USER nonroot:nonroot

ENTRYPOINT ["/product-service"]