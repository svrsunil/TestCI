FROM mysql:latest
ENV MYSQL_DATABASE=POCDB 
ENV MYSQL_ROOT_PASSWORD=root

EXPOSE 3306
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql
