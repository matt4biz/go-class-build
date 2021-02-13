FROM golang:1.15-alpine AS builder

RUN /sbin/apk update && \
	/sbin/apk --no-cache add ca-certificates git tzdata && \
	/usr/sbin/update-ca-certificates

RUN adduser -D -g '' sort
WORKDIR /home/sort

COPY go.mod /home/sort
COPY go.sum /home/sort
COPY cmd    /home/sort/cmd
COPY *.go   /home/sort

RUN CGO_ENABLED=0 go build -a -tags netgo,osusergo \
    -ldflags "-extldflags '-static' -s -w" \
    -ldflags "-X main.version=FIXME" -o sort ./cmd/sort

FROM busybox:musl

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /home/sort/sort /home/sort

USER sort
WORKDIR /home
EXPOSE 8081

ENTRYPOINT ["/home/sort"]
