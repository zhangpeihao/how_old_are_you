FROM golang:1.4
MAINTAINER Zhang Peihao "zhangpeihao@gmail.com"
LABEL Description="Demo of Face++"

RUN mkdir -p /app/how_old_are_you \
        && cd /app/how_old_are_you \
        && go get github.com/zhangpeihao/how_old_are_you
        && go build -o main .

EXPOSE 80
ENTRYPOINT ["/app/how_old_are_you/main"]
