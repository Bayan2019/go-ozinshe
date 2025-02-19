FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY ozinshe-go /bin/ozinshe-go

CMD ["ozinshe-go"]
