#FROM golang:latest
FROM balenalib/raspberry-pi-debian-golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build .
CMD ["/app/redmon"]