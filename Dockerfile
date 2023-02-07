FROM cgr.dev/chainguard/static:latest
COPY mpg /
ENTRYPOINT ["/mpg"]
