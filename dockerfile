FROM golang:1.21-alpine

WORKDIR /app

RUN mkdir ../calculator

COPY calculator/* ../calculator/

COPY calculator-api/* ./

RUN ls

RUN go mod download

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]
