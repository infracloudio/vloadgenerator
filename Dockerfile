FROM golang:alpine as builder


RUN mkdir -p /go/src/github.com/infracloudio/vloadgenerator 
ADD . /go/src/github.com/infracloudio/vloadgenerator/

WORKDIR /go/src/github.com/infracloudio/vloadgenerator
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vloadgenerator .

FROM alpine
MAINTAINER Imran Pochi <imran@infracloud.io>
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /go/src/github.com/infracloudio/vloadgenerator/vloadgenerator /app/vloadgenerator

WORKDIR /app

ENTRYPOINT ["./vloadgenerator"]
