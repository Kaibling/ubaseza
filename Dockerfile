
FROM golang:1.17.3-alpine3.15
RUN apk update && apk add make
ENTRYPOINT ["tail", "-f", "/dev/null"]