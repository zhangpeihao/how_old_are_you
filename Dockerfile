FROM golang:1.4
MAINTAINER Zhang Peihao "zhangpeihao@gmail.com"
LABEL Description="A demo of Face++"

RUN go get github.com/zhangpeihao/how_old_are_you \
    && cd $GOPATH/src/github.com/zhangpeihao/how_old_are_you \
	&& go get ./... \
    && go build -o main . \
    && cp ./main /bin/how_old_are_you

EXPOSE 80
ENTRYPOINT ["/bin/how_old_are_you"]
