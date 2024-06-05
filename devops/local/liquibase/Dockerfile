FROM liquibase/liquibase:4.2.2

USER 0:0

RUN apt-get update && apt-get install -y gettext

COPY ./docker-entrypoint-custom.sh /liquibase/docker-entrypoint-custom.sh
COPY ./init-scripts/* /dbtool/scripts/

RUN chmod +x /liquibase/docker-entrypoint-custom.sh /dbtool/scripts/*

ENTRYPOINT [ "/liquibase/docker-entrypoint-custom.sh" ]
CMD [ "update" ]
