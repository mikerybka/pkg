FROM alpine

RUN apk update
RUN apk add go

COPY . /app
WORKDIR /app

RUN go build -o /bin/server cmd/server

ENTRYPOINT ["/bin/server"]
