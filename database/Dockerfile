FROM postgres:16-alpine
COPY up.sql /docker-entrypoint-initdb.d/1.sql
EXPOSE 5432
CMD ["postgres"]