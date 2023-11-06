# module caching
FROM golang:1.21.3-alpine3.18 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# build
FROM golang:1.21.3-alpine3.18 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o math ./math

# final
FROM scratch
# copy binary file from the builder
COPY --from=builder /app/math .
EXPOSE 8989
# execute this command on "docker run"
CMD ["/math"]