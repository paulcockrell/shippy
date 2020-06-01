FROM alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

ADD vessel-service /vessel-service

CMD /wait && /vessel-service
