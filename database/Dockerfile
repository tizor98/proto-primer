FROM postgres:16-alpine
COPY up.sql /docker-entrypoint-initdb.d/1.sql
COPY triggers.sql /docker-entrypoint-initdb.d/2.sql
EXPOSE 5432
CMD ["postgres"]