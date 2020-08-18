FROM alpine
WORKDIR /root
COPY ./app /root
ENTRYPOINT ["./app"]
