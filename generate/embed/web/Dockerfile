FROM alpine:latest

RUN apk update
RUN apk add nodejs npm

COPY . /app

WORKDIR /app

RUN npm install
RUN npm run build

ENTRYPOINT ["npm", "start"]
