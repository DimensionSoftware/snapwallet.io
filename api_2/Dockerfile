FROM node:14.17.5-alpine3.14

WORKDIR /opt/app

COPY clients/ clients/
COPY error/ error/
COPY middleware/ middleware/
COPY routes/ routes/
COPY schemas/ schemas/
COPY db.js index.js package.json package-lock.json ./

RUN npm install --production=true

ENV PORT 8080
EXPOSE 8080
ENTRYPOINT [ "node", "index.js" ]
