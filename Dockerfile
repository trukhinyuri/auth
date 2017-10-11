FROM golang:1.9-alpine as builder
WORKDIR src/bitbucket.org/exonch/ch-auth
COPY . .
RUN go build -v -ldflags="-w -s" -o /bin/ch-auth

FROM alpine:latest
COPY --from=builder /bin/ch-auth /bin
# TODO: add env
CMD /bin/ch-auth