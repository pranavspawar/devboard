#syntax=docker/dockerfile:1

# Stage 1 — build the app into static files with DHI
FROM dhi.io/node:22-alpine3.24-dev AS build
WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

# Stage 2 — serve static files with lightweight DHI Nginx
FROM dhi.io/nginx:1.31.2-alpine3.24
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 4173
CMD ["nginx", "-g", "daemon off;"]
