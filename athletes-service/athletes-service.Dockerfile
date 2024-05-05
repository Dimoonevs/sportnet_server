FROM alpine:latest

RUN mkdir /app

COPY athletesApp /app

COPY /time/zoneinfo.zip /usr/local/go/lib/time/

CMD [ "/app/athletesApp" ]