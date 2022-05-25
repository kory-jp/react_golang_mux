FROM mysql:8.0.16
RUN mkdir /var/log/mysql && touch /var/log/mysql/mysqld.log
COPY conf.d/my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf