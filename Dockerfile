FROM golang:1.18-alpine

WORKDIR /api

COPY main.go ./
RUN go build -o api main.go

CMD [ "./api" ]