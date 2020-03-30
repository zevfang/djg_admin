FROM alpine:3.8

MAINTAINER zevfang <fangwei0505@126.com>
WORKDIR /home

ENV TZ=Asia/Shanghai
RUN echo "http://mirrors.aliyun.com/alpine/v3.4/main/" > /etc/apk/repositories \
    && apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

RUN apk add --no-cache ca-certificates && mkdir -p runtime/logs

COPY ./cmd/ .
VOLUME ["/hone/runtime/logs","/home/conf"]
EXPOSE 8000
ENTRYPOINT ["./djg_admin"]

