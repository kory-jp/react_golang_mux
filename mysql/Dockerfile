FROM mysql:8.0.16
ENV TZ=Asia/Tokyo
RUN mkdir /var/log/mysql
RUN touch /var/log/mysql/mysqld.log
RUN chmod 777 /var/log/mysql/mysqld.log
COPY conf.d/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./initdb.d /docker-entrypoint-initdb.d
RUN chmod 644 /etc/mysql/conf.d/my.cnf
EXPOSE 3306
CMD [ "mysqld" ]