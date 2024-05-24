# Stage 1: Build the Vue app
FROM node:22-alpine as build

WORKDIR /app

COPY package.json package-lock.json ./

# Install application dependencies
RUN npm ci

COPY . .

# Build the app
RUN npm run build

# Stage 2: Serve the Vue app with Nginx
FROM nginx:1.21-alpine

COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]