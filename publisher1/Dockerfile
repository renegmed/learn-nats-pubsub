FROM golang:1.14.0-alpine

WORKDIR src/publisher1

COPY . .

#RUN pwd
#RUN ls -als

#CMD go mod init
#CMD go mod vendor

RUN go build .

#RUN ls -als

EXPOSE 8080
ENTRYPOINT ["./publisher1"]