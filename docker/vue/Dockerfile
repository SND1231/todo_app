# ビルド環境
FROM node:13 as build-stage
WORKDIR /app
COPY ./front/package*.json ./
RUN npm install
RUN npm run vue -v
COPY ./front ./
RUN npm run build

# 本番環境
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

