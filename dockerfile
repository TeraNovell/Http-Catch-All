FROM node:lts-alpine

ENV NODE_ENV production

USER node
COPY --chown=node:node ./dist/ /var/dist/

WORKDIR /var/dist

EXPOSE 5000
ENTRYPOINT ["node", "index.js"]
