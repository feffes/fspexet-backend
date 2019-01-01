FROM golang:1.11.4

RUN go get github.com/pilu/fresh

CMD fresh