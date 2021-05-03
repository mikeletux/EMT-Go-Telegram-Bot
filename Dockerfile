FROM golang:1.16.3

WORKDIR $GOPATH/src/github.com/mikeletux/EMT-Go-Telegram-Bot

COPY . .

RUN cd cmd && \
    go build -o telegram-emt-bot && \
    mv telegram-emt-bot $GOPATH/bin

CMD ["telegram-emt-bot"]