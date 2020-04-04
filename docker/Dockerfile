ARG GO_VERSION=1.14
FROM golang:${GO_VERSION} AS go-builder

# build app
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static'" -o app .


FROM ubuntu:18.04

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get clean && \
    apt-get update && \
    apt-get install -y locales

RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

USER root
WORKDIR /root
EXPOSE 8080

COPY --from=go-builder /src/app /bin
CMD ["/bin/app", "-f", "/etc/app/config.yml"]
