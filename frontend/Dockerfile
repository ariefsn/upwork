FROM node:20-alpine AS base

ARG PORT
ARG ORIGIN
ARG PUBLIC_APP_NAME
ARG PUBLIC_API_URL
ARG PUBLIC_GOOGLE_TAG_ID

ENV PORT $PORT
ENV ORIGIN $ORIGIN
ENV PUBLIC_APP_NAME $PUBLIC_APP_NAME
ENV PUBLIC_API_URL $PUBLIC_API_URL
ENV PUBLIC_ORIGIN $ORIGIN
ENV PUBLIC_GOOGLE_TAG_ID $PUBLIC_GOOGLE_TAG_ID

# 1. Install dependencies only when needed
FROM base AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat

WORKDIR /app

# Install dependencies based on the preferred package manager
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./
RUN \
  if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm i; \
  else echo "Lockfile not found." && exit 1; \
  fi


# 2. Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/package.json ./package.json
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Heap Memory
# Please Refer https://github.com/vitejs/vite/issues/2433 or https://rollupjs.org/troubleshooting/#error-javascript-heap-out-of-memory
ENV NODE_OPTIONS=--max-old-space-size=8192

RUN yarn build

# 3. Production image, copy all the files and run
FROM base AS runner
WORKDIR /app

ENV NODE_ENV=production

RUN addgroup -g 1001 -S appgroup
RUN adduser -S appuser -u 1001

COPY --from=builder --chown=appuser:appgroup /app/package.json ./package.json
COPY --from=builder --chown=appuser:appgroup /app/node_modules ./node_modules
# needed for generating sitemap
COPY --from=builder --chown=appuser:appgroup /app/src/routes ./src/routes
COPY --from=builder --chown=appuser:appgroup /app/build ./build

USER appuser

EXPOSE $PORT

CMD [ "node", "build/index.js" ]