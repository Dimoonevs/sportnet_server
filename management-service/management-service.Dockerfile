FROM alpine:latest

RUN mkdir /app

COPY managementApp /app

COPY /time/zoneinfo.zip /usr/local/go/lib/time/

CMD [ "/app/managementApp" ]