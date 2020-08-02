FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

RUN apk add --no-cache ca-certificates git tzdata
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/vietlott .

FROM scratch

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/bin/vietlott /vietlott

ENV TZ=Asia/Ho_Chi_Minh

USER nobody:nobody

ENTRYPOINT ["/vietlott"]
