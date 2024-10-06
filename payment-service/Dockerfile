FROM golang:1.19 AS build

WORKDIR /app

COPY . ./

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /payment

FROM scratch

COPY --from=build /payment ./payment

EXPOSE 8080

CMD ["/payment"]