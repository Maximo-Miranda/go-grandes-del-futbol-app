# Stage 1: Build frontend
FROM node:24-alpine AS frontend
RUN corepack enable && corepack prepare pnpm@latest --activate
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY vite.config.ts tsconfig.json tsconfig.node.json ./
COPY resources/ resources/
RUN pnpm build

# Stage 2: Build backend
FROM golang:1.25-alpine AS backend
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend /app/public/build public/build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o grandesdelfutbol .

# Stage 3: Runtime
FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=backend /app/grandesdelfutbol .
COPY --from=frontend /app/public/build public/build
COPY resources/views resources/views
RUN mkdir -p storage/logs storage/app storage/framework/sessions storage/framework/cache
EXPOSE 3000
CMD ["./grandesdelfutbol"]
