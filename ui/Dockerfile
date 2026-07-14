# syntax=docker/dockerfile:1

##############################################################
## Stage 1 - Build the Next.js application (standalone output)
##############################################################
FROM node:20-bookworm-slim AS builder

WORKDIR /app

# Install dependencies based on the lockfile for reproducible builds
COPY package.json package-lock.json* ./
RUN npm ci

# Copy the rest of the source and build
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

##############################################################
## Stage 2 - Run the Next.js server (minimal runtime image)
##############################################################
FROM node:20-bookworm-slim AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3000
ENV HOSTNAME=0.0.0.0

# Run as an unprivileged user
RUN addgroup --system --gid 1001 nodejs \
  && adduser --system --uid 1001 nextjs

# Copy the standalone server, static assets, and public files
COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

CMD ["node", "server.js"]
