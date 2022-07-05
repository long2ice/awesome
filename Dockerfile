FROM golang AS builder
MAINTAINER long2ice "long2ice@gmail.com"
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=$GOARCH
ENV CGO_ENABLED=0
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app ./
RUN go build -o worker ./worker
RUN go build -o scheduler ./scheduler

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /build/app /
COPY --from=builder /build/worker /
COPY --from=builder /build/scheduler /
CMD ["/app"]