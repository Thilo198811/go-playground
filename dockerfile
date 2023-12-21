FROM golang:1.21

WORKDIR /app

RUN mkdir ../calculator

COPY calculator/* ../calculator/

COPY calculator-api/* ./

RUN ls

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080


CMD ["/main"]
