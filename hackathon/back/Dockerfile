FROM golang:1.21 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

COPY --from=build /app/main .

EXPOSE 8080

ENV PORT 8080

CMD ["./main"]

