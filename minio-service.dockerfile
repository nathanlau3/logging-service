# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY minioServiceApp /app

CMD [ "/app/minioServiceApp" ]