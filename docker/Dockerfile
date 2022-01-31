FROM golang:alpine

ADD main.go /go/src/hello-app/main.go
RUN cd /go/src/hello-app && \
    go mod init && \
    go install /go/src/hello-app

FROM alpine:latest

RUN apk update && \
    apk --no-cache add \
         bash \
         curl \
         dos2unix

RUN wget -P /usr/bin/ --no-check-certificate https://storage.googleapis.com/kubernetes-release/release/$(wget --no-check-certificate -O - https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x /usr/bin/kubectl

COPY --from=0 /go/bin/hello-app .
ENV PORT 8080
CMD ["./hello-app"]