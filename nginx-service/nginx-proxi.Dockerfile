# Используйте официальный образ Nginx
FROM nginx:1.19.0

# Удалите стандартный конфигурационный файл Nginx
RUN rm /etc/nginx/conf.d/default.conf

# Добавьте свой конфигурационный файл
COPY nginx.conf /etc/nginx/conf.d