FROM busybox:1.37

RUN adduser -u 1000 -D go

USER go
COPY --chown=go:go ./dist/http-catch-all-linux-amd64 /var/dist/http-catch-all-linux-amd64

WORKDIR /var/dist

EXPOSE 8080
ENTRYPOINT ["./http-catch-all-linux-amd64"]