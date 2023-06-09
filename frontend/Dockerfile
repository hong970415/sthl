ARG BUILD_ENV=local

FROM node:18-alpine as builder
ARG BUILD_ENV
WORKDIR /app
COPY next-i18next.config.js next.config.js package.json tsconfig.json yarn.lock .env.${BUILD_ENV} ./
COPY public ./public
COPY src ./src
RUN yarn install --frozen-lockfile && yarn build:${BUILD_ENV}

FROM node:18-alpine
WORKDIR /app
# https://nextjs.org/docs/pages/api-reference/next-config-js/output
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public

CMD ["node", "server.js"]