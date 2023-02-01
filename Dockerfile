FROM scratch
LABEL org.opencontainers.image.source="https://github.com/mcornick/mpg"
LABEL org.opencontainers.image.description="Generates strings that can be used as reasonably secure passwords."
LABEL org.opencontainers.image.licenses="MIT"
COPY mpg /
ENTRYPOINT ["/mpg"]
