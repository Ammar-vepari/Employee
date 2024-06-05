FROM postgres:14.8-alpine

RUN apk add gettext

COPY ./init.sh /docker-entrypoint-initdb.d/init.sh
