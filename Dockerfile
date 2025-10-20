FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/api /src/cmd/api/
RUN ls /bin

FROM alpine:3.21
COPY --from=build /bin/api /bin/api
EXPOSE 8080/tcp