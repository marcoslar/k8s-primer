FROM golang:1.19-alpine3.15 as build
WORKDIR /app
COPY main.go ./
RUN go build -o /app/HelloWorld main.go
EXPOSE 7070

FROM alpine:3.15
USER nobody:nobody
COPY --from=build /app/HelloWorld /HelloWorld

CMD ["/HelloWorld"]
