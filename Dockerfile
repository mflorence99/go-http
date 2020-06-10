FROM alpine:latest

COPY http /root

ENTRYPOINT ["/root/http"]
