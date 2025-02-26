FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /app
COPY ./../go.mod .
COPY ./../. .
# RUN go get
RUN go mod download && go mod verify
RUN GOOS="linux" GOARCH=amd64 CGO_ENABLED=0 go build -o run -a -ldflags="-s -w" -installsuffix cgo   *.go 

RUN apk --no-cache add curl ca-certificates \
 && curl https://releases.hashicorp.com/envconsul/0.6.2/envconsul_0.6.2_linux_amd64.tgz | tar -xz

RUN apk add --no-cache tzdata

FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/run .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# ENV TZ Asia/Jakarta

CMD ["/app/run"]