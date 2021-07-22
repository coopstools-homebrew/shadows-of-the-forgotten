FROM golang:1.16.5-buster AS base
WORKDIR /go/src/
COPY src/ .
RUN go build -o api .

FROM debian:buster-slim
WORKDIR /home/app
COPY --from=base /go/src/api /home/app/api
USER root:root
CMD ./api $PORT $URL_PATH_PREFIX