FROM ubuntu:14.04
MAINTAINER Jamie Patel

RUN apt-get update
RUN apt-get install -yqq graphite-web graphite-carbon libpq-dev python-psycopg2

ADD sleep.py /home/sleep.py
ADD conf/local_settings.py /etc/graphite/local_settings.py

# RUN graphite-manage syncdb

# Add config files for graphite and carbon
ADD conf/graphite-carbon /etc/default/graphite-carbon
ADD conf/carbon.conf /etc/carbon/carbon.conf
ADD conf/storage-schemas.conf /etc/carbon/storage-schemas.conf
ADD conf/storage-aggregation.conf /etc/carbon/storage-aggregation.conf

# apache
RUN apt-get -yqq install apache2 libapache2-mod-wsgi
RUN a2dissite 000-default
RUN cp /usr/share/graphite-web/apache2-graphite.conf /etc/apache2/sites-available
RUN a2ensite apache2-graphite

ADD run.sh /home/run.sh
CMD ["sh", "/home/run.sh"]
