FROM alpine:3.8

MAINTAINER zevfang <fangwei0505@126.com>
WORKDIR /home
RUN apk add --no-cache ca-certificates && mkdir -p runtime/logs
COPY ./cmd/ .
VOLUME ["/hone/runtime/logs","/home/conf"]
EXPOSE 8000
ENTRYPOINT ["./djg_admin"]

