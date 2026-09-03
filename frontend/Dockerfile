#syntax=docker/dockerfile:1

# Stage 1 — build the app into static files
FROM node:22-alpine3.24 AS build
WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

# Stage 2 — serve static files with lightweight Nginx
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 4173
CMD ["nginx", "-g", "daemon off;"]
