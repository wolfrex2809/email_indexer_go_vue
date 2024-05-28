# Stage 1: Build the Vue app
FROM node:22-alpine as build

WORKDIR /app

COPY package.json package-lock.json ./

# Install application dependencies
RUN npm ci

COPY . .
ENV NODE_ENV = production 
# Build the app
RUN npm run build

# Stage 2: Serve the Vue app with Nginx
FROM nginx:1.21-alpine

RUN apk update
RUN apk add gettext


COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/env_default.conf

ARG SERVICE_HOST=$SERVICE_HOST
ENV SERVICE_HOST=$SERVICE_HOST

RUN envsubst '$SERVICE_HOST' < /etc/nginx/conf.d/env_default.conf > /etc/nginx/conf.d/default.conf

RUN rm /etc/nginx/conf.d/env_default.conf
RUN cat /etc/nginx/conf.d/default.conf


EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]