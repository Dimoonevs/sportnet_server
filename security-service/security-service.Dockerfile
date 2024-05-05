FROM alpine:latest

RUN mkdir /app
COPY /pkg/email/template /app/template

COPY securityApp /app

CMD [ "/app/securityApp" ]