FROM golang:1.19-alpine

WORKDIR /app

ENV PYTHONUNBUFFERED=1

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk update && apk add --no-cache python3 && apk add --no-cache openssh-keygen
RUN ln -sf python3 /usr/bin/python
RUN ssh-keygen -t rsa -N "" -m PEM -f ./id.rsa
RUN ssh-keygen -f id.rsa.pub -e -m pkcs8 > id.rsa.pub.pkcs8

COPY . .
RUN go build -o /geometrics

EXPOSE 1323

CMD ["/geometrics"]
