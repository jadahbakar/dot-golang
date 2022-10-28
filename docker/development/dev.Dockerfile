# Please keep up to date with the new-version of Golang docker for builder
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git && apk --no-cache add curl
WORKDIR /app 
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
CMD air