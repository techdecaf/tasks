FROM alpine:3.14.0
COPY "./build/linux/tasks" "/usr/local/bin"
USER go