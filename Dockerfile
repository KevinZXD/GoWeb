FROM 1067892503/gosource

RUN mkdir -p /data0/htdocs/GoWeb

ADD GoWeb.tgz /data0/htdocs/GoWeb

ADD run.sh  /

CMD ["sh","run.sh"]