FROM nginx:latest

RUN rm /etc/nginx/conf.d/default.conf

COPY ./modules /etc/nginx/modules
COPY ./nginx.conf /etc/nginx