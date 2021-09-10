FROM golang:1.16

RUN apt-get update && apt-get install -y npm=5.8.0+ds6-4+deb10u2

RUN npm install -g yarn@1.22.10

RUN mkdir -p /app/web-client

COPY ./development/frontend.sh /frontend.sh

RUN chmod +x /frontend.sh

ENTRYPOINT ["/frontend.sh"]
