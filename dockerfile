FROM busybox:1.37

RUN adduser -u 1000 -D go

USER go
COPY --chown=go:go ./dist/ /var/dist/

WORKDIR /var/dist

EXPOSE 8080
ENTRYPOINT ["./http-catch-all"]