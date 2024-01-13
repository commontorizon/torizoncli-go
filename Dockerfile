FROM scratch
COPY torizoncli /usr/bin/torizoncli
ENTRYPOINT ["/usr/bin/torizoncli"]
