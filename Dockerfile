FROM golang:1.21 as build
RUN apt update
RUN apt-get install -y libpcap-dev
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go mod vendor
RUN go build -o sniffer main.go
FROM scratch
COPY --from=build /app/sniffer /sniffer
ENTRYPOINT ["/sniffer"]