FROM alpine:latest

RUN mkdir /app

COPY subscriptionApp /app

CMD [ "/app/subscriptionApp" ]