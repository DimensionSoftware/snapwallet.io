FROM node:14.17.5-alpine3.14

WORKDIR /opt/app

COPY package.json package-lock.json ./
RUN npm install --production=true

COPY clients/ clients/
COPY error/ error/
COPY middleware/ middleware/
COPY routes/ routes/
COPY schemas/ schemas/
COPY util/ util/
COPY db.js dbevents.js index.js jobpublisher.js ./


ENV PORT 8080
EXPOSE 8080
ENTRYPOINT [ "node", "index.js" ]
