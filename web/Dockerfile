FROM node:lts-alpine AS web

WORKDIR /app
COPY . /app
RUN cd /app && npm install && npm run build

FROM nginx:stable-alpine

COPY --from=web /app/dist /usr/share/nginx/html
RUN rm -f /etc/nginx/conf.d/*
COPY /nginx/nginx.conf /etc/nginx/conf.d/my.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
