###############################################################################
#                                  BUILD
###############################################################################

FROM golang as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

###############################################################################
#                                INSTALLATION
###############################################################################

FROM ubuntu:22.04

ENV WORKDIR /app
ADD hack $WORKDIR/hack/
ADD manifest/config/config.yaml.example $WORKDIR/manifest/config/config.yaml
ADD resource $WORKDIR/resource/
ADD storage $WORKDIR/storage/

COPY --from=builder /app/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

RUN apt-get update -y \
    && DEBIAN_FRONTEND=noninteractive TZ=Asia/Shanghai apt-get install -y tzdata ca-certificates

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
