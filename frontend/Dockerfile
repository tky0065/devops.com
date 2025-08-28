### Frontend multi-stage Dockerfile: build with Node, serve with nginx
FROM node:20-alpine AS builder

WORKDIR /app

# Install deps
COPY package.json package-lock.json* ./
RUN npm ci --silent || npm install --silent

# Copy sources and build
COPY . .
RUN npm run build

FROM nginx:stable-alpine

# Copy built assets
COPY --from=builder /app/dist /usr/share/nginx/html

# Add nginx config (will proxy /api to backend)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
