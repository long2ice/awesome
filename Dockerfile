FROM node as web-builder
ENV REACT_APP_API_URL=/api
RUN mkdir -p /src
WORKDIR /src
RUN git clone https://github.com/long2ice/awesome-web.git
RUN cd awesome-web && npm i && npm run build

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
COPY --from=web-builder /src/awesome-web/build /build/static
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