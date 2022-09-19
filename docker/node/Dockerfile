FROM node:latest as builder

WORKDIR /app

COPY ./front/package.json package.json
COPY ./front/yarn.lock yarn.lock

RUN yarn install

COPY ./front /app

RUN NODE_ENV=production yarn build


FROM node:latest as runner

WORKDIR /app

ENV NODE_ENV production

COPY --from=builder /app/next.config.js next.config.js
COPY --from=builder /app/.next .next
COPY --from=builder /app/.env.production .env
COPY --from=builder /app/public public
COPY --from=builder /app/node_modules node_modules
COPY --from=builder /app/package.json package.json
COPY --from=builder /app/yarn.lock yarn.lock

CMD ["yarn", "start"]