FROM debian

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

ADD email-service /email-service

CMD /wait && /email-service
