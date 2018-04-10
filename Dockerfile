FROM alpine:3.1
MAINTAINER Anubhav Mishra <anubhavmishra@me.com>
ADD build/linux/amd64/key-count /usr/bin/key-count
ENTRYPOINT ["key-count"]