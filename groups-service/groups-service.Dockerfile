FROM alpine:latest

RUN mkdir /app

COPY groupsApp /app

CMD [ "/app/groupsApp" ]