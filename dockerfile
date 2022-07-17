FROM node:lts-alpine

ENV NODE_ENV production

COPY ./dist/ /var/dist/

WORKDIR /var/dist

USER node

EXPOSE 5000
ENTRYPOINT ["node", "index.js"]
