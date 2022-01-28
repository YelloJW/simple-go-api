# syntax=docker/dockerfile:1
FROM golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN apk add --no-cache build-base &&\
    make build

FROM golang:1.17-alpine
WORKDIR /app
COPY --from=builder app/main .
EXPOSE 8081
CMD [ "/app/main" ] 

